package main

import (
	"path/filepath"

	"github.com/alecthomas/kong"

	"github.com/ibrt/golang-lib/consolez"
	"github.com/ibrt/golang-lib/devz"
	"github.com/ibrt/golang-lib/errorz"
)

func main() {
	defer consolez.DefaultCLI.Recover(true)

	cli := &CLI{}
	k := kong.Parse(cli)
	errorz.MaybeMustWrap(k.Run())
}

// CLI describes the command line interface.
type CLI struct {
	Test     *TestCmd     `cmd:"" help:"Run tests."`
	Validate *ValidateCmd `cmd:"" help:"Build, lint, and test everything."`
}

// AfterApply is a Kong hook.
func (c *CLI) AfterApply(k *kong.Context) error {
	consolez.DefaultCLI.Tool("golang-lib", k)
	return nil
}

// TestCmd implements a command.
type TestCmd struct {
	Packages     []string `arg:"" optional:"" help:"Select specific packages."`
	IgnoreCache  bool     `flag:"" help:"Force running all tests, even if the results are cached."`
	OpenCoverage bool     `flag:"" help:"Open coverage results after running."`
}

// Run the command.
func (c *TestCmd) Run() error {
	devz.MustRunGoTests(&devz.GoTestsParams{
		AllPackages:      []string{"./..."},
		SelectedPackages: c.Packages,
		IgnoreCache:      c.IgnoreCache,
		CoverageDirPath:  filepath.Join(".build", "coverage"),
		OpenCoverage:     c.OpenCoverage,
		PrintNotices:     true,
		PrintCommands:    true,
	})

	return nil
}

// ValidateCmd implements a command.
type ValidateCmd struct {
	// intentionally empty
}

// Run the command.
func (c *ValidateCmd) Run() error {
	devz.MustRunGoChecks(&devz.GoChecksParams{
		AllPackages:   []string{"./..."},
		PrintNotices:  true,
		PrintCommands: true,
	})

	devz.MustRunGoTests(&devz.GoTestsParams{
		AllPackages:     []string{"./..."},
		IgnoreCache:     true,
		CoverageDirPath: filepath.Join(".build", "coverage"),
		PrintNotices:    true,
		PrintCommands:   true,
	})

	return nil
}
