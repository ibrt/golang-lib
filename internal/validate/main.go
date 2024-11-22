package main

import (
	"path/filepath"

	"github.com/ibrt/golang-lib/consolez"
	"github.com/ibrt/golang-lib/devz"
)

func main() {
	consolez.DefaultCLI.Banner("golang-lib", "Validate")

	devz.MustRunGoChecks(&devz.GoChecksParams{
		Packages:      []string{"./..."},
		PrintNotices:  true,
		PrintCommands: true,
	})

	devz.MustRunGoTests(&devz.GoTestsParams{
		Packages:        []string{"./..."},
		CoverageDirPath: filepath.Join(".build", "coverage"),
		OpenCoverage:    true,
		PrintNotices:    true,
		PrintCommands:   true,
	})
}
