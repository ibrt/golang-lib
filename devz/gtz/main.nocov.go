package main

import (
	"os"

	"github.com/ibrt/golang-lib/consolez"
	"github.com/ibrt/golang-lib/devz"
	"github.com/ibrt/golang-lib/errorz"
)

func main() {
	defer consolez.DefaultCLI.Recover(false)

	if len(os.Args) < 2 {
		panic(errorz.Errorf("usage: gtz <go-tool-key> [args...]"))
	}

	devz.MustLookupGoTool(os.Args[1]).
		GetCommand().
		AddParams(os.Args[2:]...).
		MustExec()
}
