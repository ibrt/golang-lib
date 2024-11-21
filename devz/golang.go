package devz

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ibrt/golang-lib/consolez"
	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/jsonz"
	"github.com/ibrt/golang-lib/memz"
	"github.com/ibrt/golang-lib/shellz"
	"github.com/ibrt/golang-lib/stringz"
)

// GoTool describes a Go tool.
type GoTool struct {
	m              *sync.Mutex
	pkg            string
	path           string
	defaultVersion string
	currentVersion string
}

// Known Go tools.
var (
	GoToolGoCov       = NewGoTool("github.com/axw/gocov", "gocov", "v1.2.1")
	GoToolGoCovHTML   = NewGoTool("github.com/matm/gocov-html", "cmd/gocov-html", "v1.4.0")
	GoToolGolint      = NewGoTool("golang.org/x/lint", "golint", "v0.0.0-20210508222113-6edffad5e616")
	GoToolMockGen     = NewGoTool("go.uber.org/mock", "/mockgen", "v0.5.0")
	GoToolStaticCheck = NewGoTool("honnef.co/go/tools", "cmd/staticcheck", "2024.1.1")
)

// NewGoTool initializes a new Go tool.
func NewGoTool(pkg, path, defaultVersion string) *GoTool {
	if path != "" {
		path = stringz.EnsurePrefix(path, "/")
	}

	return &GoTool{
		m:              &sync.Mutex{},
		pkg:            strings.TrimSuffix(pkg, "/"),
		path:           path,
		defaultVersion: defaultVersion,
		currentVersion: "",
	}
}

// GetVersion returns the current Go tool defaultVersion (either pinned, or from go.mod).
func (t *GoTool) GetVersion() string {
	t.m.Lock()
	defer t.m.Unlock()
	return t.getVersion()
}

// GetPackage returns the Go package.
func (t *GoTool) GetPackage() string {
	t.m.Lock()
	defer t.m.Unlock()
	return t.pkg
}

// GetArgument returns an argument string suitable for Go run.
func (t *GoTool) GetArgument() string {
	t.m.Lock()
	defer t.m.Unlock()
	return fmt.Sprintf("%v%v@%v", t.pkg, t.path, t.getVersion())
}

// GetCommand returns a *shellz.Command for this Go tool.
func (t *GoTool) GetCommand() *shellz.Command {
	return shellz.NewCommand("go", "run").AddParams(t.GetArgument())
}

// MustRun runs the Go tool.
func (t *GoTool) MustRun(params ...string) {
	t.GetCommand().AddParams(params...).MustRun()
}

func (t *GoTool) getVersion() string {
	if t.currentVersion != "" {
		return t.currentVersion
	}

	func() {
		defer func() { recover() }()
		out := shellz.NewCommand("go", "list", "-m", t.pkg).SetEcho(false).MustOutputString()
		t.currentVersion = strings.TrimSpace(strings.TrimPrefix(out, t.pkg))
	}()

	if t.currentVersion == "" {
		if t.defaultVersion != "" {
			t.currentVersion = t.defaultVersion
		} else {
			t.currentVersion = "latest"
		}
	}

	return t.currentVersion
}

// GoChecksParams describes the parameters for running Go checks.
type GoChecksParams struct {
	Packages      []string
	BuildTags     []string
	PrintNotices  bool
	PrintCommands bool
}

// MustRunGoChecks runs a set of Go checks in the current working directory.
func MustRunGoChecks(params *GoChecksParams) {
	errorz.Assertf(len(params.Packages) > 0, "missing packages")

	if params.PrintNotices {
		consolez.DefaultCLI.Notice("go-checks", "preparing...")
	}

	shellz.NewCommand("go", "mod", "tidy").
		SetEcho(params.PrintCommands).
		MustRun()

	shellz.NewCommand("go", "generate").
		AddParams(params.Packages...).
		SetEcho(params.PrintCommands).
		MustRun()

	shellz.NewCommand("go", "fmt").
		AddParams(params.Packages...).
		SetEcho(params.PrintCommands).
		MustRun()

	if params.PrintNotices {
		consolez.DefaultCLI.Notice("go-checks", "building...")
	}

	shellz.NewCommand("go", "build", "-v").
		AddParamsIfTrue(len(params.BuildTags) > 0, fmt.Sprintf("-tags=%v", strings.Join(params.BuildTags, ","))).
		AddParams(params.Packages...).
		SetEcho(params.PrintCommands).
		MustRun()

	if params.PrintNotices {
		consolez.DefaultCLI.Notice("go-checks", "linting...")
	}

	GoToolGolint.
		GetCommand().
		AddParams("-set_exit_status").
		AddParams(params.Packages...).
		SetEcho(params.PrintCommands).
		MustRun()

	shellz.NewCommand("go", "vet").
		AddParams(params.Packages...).
		SetEcho(params.PrintCommands).
		MustRun()

	GoToolStaticCheck.
		GetCommand().
		AddParams(params.Packages...).
		SetEcho(params.PrintCommands).
		MustRun()

	shellz.NewCommand("go", "mod", "tidy").
		SetEcho(params.PrintCommands).
		MustRun()
}

// GoTestsParams describes the parameters for running Go tests.
type GoTestsParams struct {
	Packages        []string
	BuildTags       []string
	TestRegexp      string
	IgnoreCache     bool
	Verbose         *bool
	CoverageDirPath string
	OpenCoverage    bool
	PrintNotices    bool
	PrintCommands   bool
}

// MustRunGoTests runs a set of Go tests in the current working directory.
func MustRunGoTests(params *GoTestsParams) {
	errorz.Assertf(len(params.Packages) > 0, "missing packages")
	errorz.Assertf(params.CoverageDirPath != "", "missing coverage dir path")

	if params.PrintNotices {
		consolez.DefaultCLI.Notice("go-tests", "preparing coverage directory...")
	}

	filez.MustPrepareDir(params.CoverageDirPath, 0777)

	if params.PrintNotices {
		consolez.DefaultCLI.Notice("go-tests", "generating Go code...")
	}

	shellz.NewCommand("go", "generate").
		AddParams(params.Packages...).
		SetEcho(params.PrintCommands).
		MustRun()

	cmd := shellz.NewCommand("go", "test").
		AddParams("-trimpath", "-race", "-failfast", "-shuffle=on").
		AddParams("-covermode=atomic", fmt.Sprintf("-coverprofile=%v", filepath.Join(params.CoverageDirPath, "coverage.out")))

	if params.IgnoreCache {
		cmd = cmd.AddParams("-count=1")
	}

	if params.TestRegexp != "" {
		cmd = cmd.AddParams(fmt.Sprintf("-run=%v", params.TestRegexp))
	}

	if memz.ValNilToZero(params.Verbose) ||
		(params.Verbose == nil && len(params.Packages) == 1 && !strings.HasSuffix(params.Packages[0], "...")) {
		cmd = cmd.AddParams("-v")
	}

	if params.PrintNotices {
		consolez.DefaultCLI.Notice("go-tests", "running tests...")
	}

	p := consolez.NewGoTestPrinter()
	cmd.AddParams(params.Packages...).MustLines(p.PrintLine)
	p.PrintDone()

	if params.PrintNotices {
		consolez.DefaultCLI.Notice("go-tests", "processing coverage...")
	}

	coverageJSON := processGoCoverage(params)
	consolez.NewCoveragePrinter().Print(jsonz.MustUnmarshal[*consolez.Coverage](coverageJSON))

	if params.OpenCoverage {
		if params.PrintNotices {
			consolez.DefaultCLI.Notice("go-tests", "opening coverage...")
		}

		openGoCoverage(params, coverageJSON)
	}
}

func processGoCoverage(params *GoTestsParams) []byte {
	coverageOutLines := memz.FilterSlice(
		strings.Split(filez.MustReadFileString(filepath.Join(params.CoverageDirPath, "coverage.out")), "\n"),
		func(l string) bool {
			return !strings.Contains(l, ".gen.go:") && !strings.Contains(l, "_generated.go:")
		})

	filez.MustWriteFileString(
		filepath.Join(params.CoverageDirPath, "coverage.out"),
		0777, 0666,
		strings.Join(coverageOutLines, "\n"))

	coverageJSON := GoToolGoCov.
		GetCommand().
		AddParams("convert", filepath.Join(params.CoverageDirPath, "coverage.out")).
		MustOutput()

	filez.MustWriteFile(
		filepath.Join(params.CoverageDirPath, "coverage.json"),
		0777, 0666,
		coverageJSON)

	return coverageJSON
}

func openGoCoverage(params *GoTestsParams, coverageJSON []byte) {
	coverageHTML := GoToolGoCovHTML.
		GetCommand().
		AddParams("-t", "golang").
		SetIn(bytes.NewReader(coverageJSON)).
		MustOutput()

	filez.MustWriteFile(
		filepath.Join(params.CoverageDirPath, "coverage.html"),
		0777, 0666,
		coverageHTML)

	shellz.NewCommand("open").
		AddParams(filepath.Join(params.CoverageDirPath, "coverage.html")).
		MustRun()
}
