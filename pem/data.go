package pem

import (
	"path/filepath"
	"runtime"
)

var basepath string

func init() {
	_, thisFile, _, _ := runtime.Caller(0)
	basepath = filepath.Base(thisFile)
}

func LoadPath(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}
	return filepath.Join(basepath, rel)
}
