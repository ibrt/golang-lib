package filez

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ibrt/golang-lib/errorz"
	"github.com/ibrt/golang-lib/stringz"
)

// MustAbs is like filepath.Abs, but panics on error.
func MustAbs(path string) string {
	path, err := filepath.Abs(path)
	errorz.MaybeMustWrap(err)
	return path
}

// MustRel is like filepath.Rel but panics on error.
func MustRel(src, dst string) string {
	path, err := filepath.Rel(MustAbs(src), MustAbs(dst))
	errorz.MaybeMustWrap(err)
	return path
}

// MaybeMustRelIfChild is like MustRel but only gets applied if "filePath" is an absolute path and a child of "parent".
func MaybeMustRelIfChild(filePath, parentDirPath string) string {
	parentDirPath = MustAbs(parentDirPath)
	parentPrefix := stringz.EnsureSuffix(parentDirPath, string(os.PathSeparator))

	if filepath.IsAbs(filePath) && strings.HasPrefix(filePath, parentPrefix) {
		return MustRel(parentDirPath, filePath)
	}

	return filePath
}

// MustGetwd is like os.Getwd, but panics on error.
func MustGetwd() string {
	wd, err := os.Getwd()
	errorz.MaybeMustWrap(err)
	return wd
}

// MustChdir is like os.Chdir, but panics on error.
func MustChdir(wd string) string {
	errorz.MaybeMustWrap(os.Chdir(wd))
	return wd
}

// MustUserHomeDir is like os.UserHomeDir, but panics on error.
func MustUserHomeDir() string {
	dirPath, err := os.UserHomeDir()
	errorz.MaybeMustWrap(err)
	return dirPath
}
