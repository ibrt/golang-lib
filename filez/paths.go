package filez

import (
	"path/filepath"
	"strings"
)

// MustIsChild returns true if "childPath" is lexically determined to be a child of "parentPath". Panics on error.
func MustIsChild(parentPath, childPath string) bool {
	rel := MustRel(parentPath, childPath)
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

// MustRelForDisplay converts "path" to relative if (1) it is an absolute path, and (2) it is a child of the current
// working directory. It returns "path" cleaned otherwise. Panics on error.
func MustRelForDisplay(path string) string {
	if wd := MustGetwd(); filepath.IsAbs(path) && MustIsChild(wd, path) {
		return MustRel(wd, path)
	}

	return filepath.Clean(path)
}
