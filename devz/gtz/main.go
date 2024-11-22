package main

import (
	"os"

	"github.com/ibrt/golang-lib/devz"
)

var (
	knownTools = map[string]*devz.GoTool{
		"mock-gen": devz.GoToolMockGen,
	}
)

func main() {
	knownTools[os.Args[1]].GetCommand().AddParams(os.Args[2:]...).MustExec()
}
