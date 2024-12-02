package filez

import (
	"os"
	"path/filepath"

	"github.com/ibrt/golang-lib/errorz"
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

// MustRemoveAll is like os.RemoveAll, but panics on error.
func MustRemoveAll(path string) {
	errorz.MaybeMustWrap(os.RemoveAll(path))
}
