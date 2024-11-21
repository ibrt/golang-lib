package gtz

import (
	"fmt"
	"strings"
	"sync"

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
	GoToolGoLint      = NewGoTool("golang.org/x/lint", "golint", "v0.0.0-20210508222113-6edffad5e616")
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
