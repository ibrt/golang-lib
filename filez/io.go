package filez

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/ibrt/golang-lib/errorz"
)

// MustReadFile reads a file, panics on error.
func MustReadFile(filePath string) []byte {
	buf, err := os.ReadFile(filePath)
	errorz.MaybeMustWrap(err)
	return buf
}

// MustReadFileString reads a file, panics on error.
func MustReadFileString(filePath string) string {
	return string(MustReadFile(filePath))
}

// MustWriteFile creates a file with the given mode and contents, also ensuring the containing folder exists.
func MustWriteFile(filePath string, dirMode os.FileMode, fileMode os.FileMode, contents []byte) string {
	errorz.MaybeMustWrap(os.MkdirAll(filepath.Dir(filePath), dirMode))
	errorz.MaybeMustWrap(os.WriteFile(filePath, contents, fileMode))
	return filePath
}

// MustWriteFileString creates a file with the given mode and contents, also ensuring the containing folder exists.
func MustWriteFileString(filePath string, dirMode os.FileMode, fileMode os.FileMode, contents string) string {
	return MustWriteFile(filePath, dirMode, fileMode, []byte(contents))
}

// MustCreateTempFile creates a temporary file with the given contents.
func MustCreateTempFile(contents []byte) string {
	fd, err := os.CreateTemp("", "golang-lib-")
	errorz.MaybeMustWrap(err)
	defer errorz.MustClose(fd)

	_, err = io.Copy(fd, bytes.NewReader(contents))
	errorz.MaybeMustWrap(err)
	return fd.Name()
}

// MustCreateTempFileString creates a temporary file with the given contents.
func MustCreateTempFileString(contents string) string {
	return MustCreateTempFile([]byte(contents))
}

// MustCreateTempDir is like os.MkdirTemp, but panics on error.
func MustCreateTempDir() string {
	dirPath, err := os.MkdirTemp("", "golang-lib-")
	errorz.MaybeMustWrap(err)
	return dirPath
}

// MustPrepareDir deletes the given directory and its contents (if present) and recreates it.
func MustPrepareDir(dirPath string, dirMode os.FileMode) {
	errorz.MaybeMustWrap(os.RemoveAll(dirPath))
	errorz.MaybeMustWrap(os.MkdirAll(dirPath, dirMode))
}

// MustCheckPathExists checks if the given path exists, panics on errors other than os.ErrNotExist.
func MustCheckPathExists(fileOrDirPath string) bool {
	if _, err := os.Stat(fileOrDirPath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		errorz.MustWrap(err)
	}
	return true
}

// MustCheckFileExists checks if the given path exists and is a regular file, panics on errors other than os.ErrNotExist.
func MustCheckFileExists(fileOrDirPath string) bool {
	stat, err := os.Stat(fileOrDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		errorz.MustWrap(err)
	}

	return stat.Mode().IsRegular()
}
