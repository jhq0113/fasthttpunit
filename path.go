package fasthttpunit

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	binPath string
)

func init() {
	loadBinPath()
}

func loadBinPath() {
	if len(os.Args) > 2 && strings.HasPrefix(os.Args[2], "-test.") {
		if dir, _ := os.Getwd(); len(dir) > 0 {
			binPath = dir
			return
		}
	}

	binPath = filepath.Dir(os.Args[0])
}

func BinPath() string {
	return binPath
}
