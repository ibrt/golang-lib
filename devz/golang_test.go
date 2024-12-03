package devz_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/ibrt/golang-lib/devz"
	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/filez"
	"github.com/ibrt/golang-lib/fixturez"
	"github.com/ibrt/golang-lib/shellz"
	"github.com/ibrt/golang-lib/shellz/tshellz"
)

type GoSuite struct {
	// intentionally empty
}

func TestGoSuite(t *testing.T) {
	fixturez.RunSuite(t, &GoSuite{})
}

func (*GoSuite) TestMustLookupGoTool(g *WithT) {
	g.Expect(devz.MustLookupGoTool("go-cov")).To(Equal(devz.GoToolGoCov))
	g.Expect(devz.MustLookupGoTool("go-cov-html")).To(Equal(devz.GoToolGoCovHTML))
	g.Expect(devz.MustLookupGoTool("golint")).To(Equal(devz.GoToolGolint))
	g.Expect(devz.MustLookupGoTool("mock-gen")).To(Equal(devz.GoToolMockGen))
	g.Expect(devz.MustLookupGoTool("static-check")).To(Equal(devz.GoToolStaticCheck))
	g.Expect(func() { devz.MustLookupGoTool("unknown") }).To(PanicWith(MatchError("unknown go tool: unknown")))
}

func (*GoSuite) TestGoTool(g *WithT) {
	devz.GoToolGolint.MustRun(".")

	g.Expect(devz.NewGoTool("a", "b", "c").GetPackage()).To(Equal("a"))
	g.Expect(devz.NewGoTool("a", "b", "c").GetArgument()).To(Equal("a/b@c"))
	g.Expect(devz.NewGoTool("a", "b", "c").GetVersion()).To(Equal("c"))
	g.Expect(devz.NewGoTool("github.com/axw/gocov", "", "unused").GetVersion()).To(Equal("v1.2.1"))

	gt := devz.NewGoTool("a", "b", "")
	g.Expect(gt.GetVersion()).To(Equal("latest"))
	g.Expect(gt.GetVersion()).To(Equal("latest"))
}

func (*GoSuite) TestRunGoChecks(g *WithT, ctrl *gomock.Controller) {
	devz.GoToolGolint.GetVersion()      // warm up
	devz.GoToolStaticCheck.GetVersion() // warm up

	m := tshellz.NewMockExecutor(ctrl)
	shellz.DefaultExecutor = m
	defer shellz.RestoreDefaultExecutor()

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "mod", "tidy"})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "generate", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "fmt", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "build", "-v", "-tags=t1,t2", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "run", devz.GoToolGolint.GetArgument(), "-set_exit_status", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "vet", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "run", devz.GoToolStaticCheck.GetArgument(), "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "mod", "tidy"})
		})).
		Times(1).
		Return(nil)

	fixturez.MustBeginOutputCapture(fixturez.OutputSetupStandard, fixturez.GetOutputSetupColor(true), fixturez.OutputSetupTable)
	defer fixturez.ResetOutputCapture()

	devz.MustRunGoChecks(&devz.GoChecksParams{
		AllPackages: []string{"./..."},
		BuildTags:   []string{"t1", "t2"},
	})

	outBuf, errBuf := fixturez.MustEndOutputCapture()
	g.Expect(outBuf).To(Equal(strings.Join([]string{
		"[...............go-checks] preparing...",
		"🏃 go mod tidy",
		"🏃 go generate ./...",
		"🏃 go fmt ./...",
		"[...............go-checks] building...",
		"🏃 go build -v -tags=t1,t2 ./...",
		"[...............go-checks] linting...",
		fmt.Sprintf("🏃 go run %v -set_exit_status ./...", devz.GoToolGolint.GetArgument()),
		"🏃 go vet ./...",
		fmt.Sprintf("🏃 go run %v ./...", devz.GoToolStaticCheck.GetArgument()),
		"🏃 go mod tidy",
		"",
	}, "\n")))
	g.Expect(errBuf).To(BeEmpty())
}

func (*GoSuite) TestRunGoTests_SelectedPackages(g *WithT, ctrl *gomock.Controller) {
	devz.GoToolGoCov.GetVersion()     // warm up
	devz.GoToolGoCovHTML.GetVersion() // warm up

	m := tshellz.NewMockExecutor(ctrl)
	shellz.DefaultExecutor = m
	defer shellz.RestoreDefaultExecutor()

	dirPath := filez.MustCreateTempDir()
	defer filez.MustRemoveAll(dirPath)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "generate", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdStart(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			isMatch := reflect.DeepEqual(c.Args, []string{
				"go", "test",
				"-trimpath", "-race", "-failfast", "-shuffle=on", "-covermode=atomic",
				fmt.Sprintf("-coverprofile=%v", filepath.Join(dirPath, "coverage.out")),
				"-count=1", "-run=^test$", "-v",
				"./package",
			})

			if isMatch {
				errorz.MaybeMustWrap(c.Stdout.(*os.File).Close())
				errorz.MaybeMustWrap(c.Stderr.(*os.File).Close())
			}

			return isMatch
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdWait(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			isMatch := reflect.DeepEqual(c.Args, []string{
				"go", "test",
				"-trimpath", "-race", "-failfast", "-shuffle=on", "-covermode=atomic",
				fmt.Sprintf("-coverprofile=%v", filepath.Join(dirPath, "coverage.out")),
				"-count=1", "-run=^test$", "-v",
				"./package",
			})

			if isMatch {
				filez.MustWriteFileString(filepath.Join(dirPath, "coverage.out"), 0777, 0666, "")
			}

			return isMatch
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdOutput(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{
				"go", "run",
				devz.GoToolGoCov.GetArgument(),
				"convert",
				filepath.Join(dirPath, "coverage.out")})
		})).
		Times(1).
		Return([]byte("{}"), nil)

	m.EXPECT().ExecCmdOutput(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{
				"go", "run",
				devz.GoToolGoCovHTML.GetArgument(),
				"-t", "golang"})
		})).
		Times(1).
		Return([]byte("<html></html>"), nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"open", filepath.Join(dirPath, "coverage.html")})
		})).
		Times(1).
		Return(nil)

	fixturez.MustBeginOutputCapture(fixturez.OutputSetupStandard, fixturez.GetOutputSetupColor(true), fixturez.OutputSetupTable)
	defer fixturez.ResetOutputCapture()

	devz.MustRunGoTests(&devz.GoTestsParams{
		AllPackages:      []string{"./..."},
		SelectedPackages: []string{"./package"},
		BuildTags:        []string{"t1", "t2"},
		TestRegexp:       "^test$",
		IgnoreCache:      true,
		Verbose:          nil,
		CoverageDirPath:  dirPath,
		OpenCoverage:     true,
	})

	outBuf, errBuf := fixturez.MustEndOutputCapture()
	g.Expect(outBuf).To(Equal(strings.Join([]string{
		"[................go-tests] preparing coverage directory...",
		"[................go-tests] generating Go code...",
		"🏃 go generate ./...",
		"[................go-tests] running tests...",
		fmt.Sprintf("🏃 go test -trimpath -race -failfast -shuffle=on -covermode=atomic -coverprofile=%v/coverage.out -count=1 -run=^test$ -v ./package", dirPath),
		"DONE    [SKIP: 0, PASS: 0]                                           0s        ",
		"[................go-tests] processing coverage...",
		"DONE    [LOWC: 0, MEDC: 0, HIGC: 0]                                  100.0% [0/0]",
		"[................go-tests] opening coverage...",
		fmt.Sprintf("🏃 open %v/coverage.html", dirPath),
		"",
	}, "\n")))
	g.Expect(errBuf).To(BeEmpty())
}

func (*GoSuite) TestRunGoTests_AllPackages(g *WithT, ctrl *gomock.Controller) {
	devz.GoToolGoCov.GetVersion()     // warm up
	devz.GoToolGoCovHTML.GetVersion() // warm up

	m := tshellz.NewMockExecutor(ctrl)
	shellz.DefaultExecutor = m
	defer shellz.RestoreDefaultExecutor()

	dirPath := filez.MustCreateTempDir()
	defer filez.MustRemoveAll(dirPath)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"go", "generate", "./..."})
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdStart(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			isMatch := reflect.DeepEqual(c.Args, []string{
				"go", "test",
				"-trimpath", "-race", "-failfast", "-shuffle=on", "-covermode=atomic",
				fmt.Sprintf("-coverprofile=%v", filepath.Join(dirPath, "coverage.out")),
				"-count=1", "-run=^test$",
				"./...",
			})

			if isMatch {
				errorz.MaybeMustWrap(c.Stdout.(*os.File).Close())
				errorz.MaybeMustWrap(c.Stderr.(*os.File).Close())
			}

			return isMatch
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdWait(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			isMatch := reflect.DeepEqual(c.Args, []string{
				"go", "test",
				"-trimpath", "-race", "-failfast", "-shuffle=on", "-covermode=atomic",
				fmt.Sprintf("-coverprofile=%v", filepath.Join(dirPath, "coverage.out")),
				"-count=1", "-run=^test$",
				"./...",
			})

			if isMatch {
				filez.MustWriteFileString(filepath.Join(dirPath, "coverage.out"), 0777, 0666, "")
			}

			return isMatch
		})).
		Times(1).
		Return(nil)

	m.EXPECT().ExecCmdOutput(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{
				"go", "run",
				devz.GoToolGoCov.GetArgument(),
				"convert",
				filepath.Join(dirPath, "coverage.out")})
		})).
		Times(1).
		Return([]byte("{}"), nil)

	m.EXPECT().ExecCmdOutput(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{
				"go", "run",
				devz.GoToolGoCovHTML.GetArgument(),
				"-t", "golang"})
		})).
		Times(1).
		Return([]byte("<html></html>"), nil)

	m.EXPECT().ExecCmdRun(
		gomock.Any(),
		gomock.Cond(func(c *exec.Cmd) bool {
			return reflect.DeepEqual(c.Args, []string{"open", filepath.Join(dirPath, "coverage.html")})
		})).
		Times(1).
		Return(nil)

	fixturez.MustBeginOutputCapture(fixturez.OutputSetupStandard, fixturez.GetOutputSetupColor(true), fixturez.OutputSetupTable)
	defer fixturez.ResetOutputCapture()

	devz.MustRunGoTests(&devz.GoTestsParams{
		AllPackages:      []string{"./..."},
		SelectedPackages: nil,
		BuildTags:        []string{"t1", "t2"},
		TestRegexp:       "^test$",
		IgnoreCache:      true,
		Verbose:          nil,
		CoverageDirPath:  dirPath,
		OpenCoverage:     true,
	})

	outBuf, errBuf := fixturez.MustEndOutputCapture()
	g.Expect(outBuf).To(Equal(strings.Join([]string{
		"[................go-tests] preparing coverage directory...",
		"[................go-tests] generating Go code...",
		"🏃 go generate ./...",
		"[................go-tests] running tests...",
		fmt.Sprintf("🏃 go test -trimpath -race -failfast -shuffle=on -covermode=atomic -coverprofile=%v/coverage.out -count=1 -run=^test$ ./...", dirPath),
		"DONE    [SKIP: 0, PASS: 0]                                           0s        ",
		"[................go-tests] processing coverage...",
		"DONE    [LOWC: 0, MEDC: 0, HIGC: 0]                                  100.0% [0/0]",
		"[................go-tests] opening coverage...",
		fmt.Sprintf("🏃 open %v/coverage.html", dirPath),
		"",
	}, "\n")))
	g.Expect(errBuf).To(BeEmpty())
}
