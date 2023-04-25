package file

import (
	"os"
	"path/filepath"
)

// Mkdir Create the DIRECTORY(ies), if they do not already exist
// parents no error if existing, make parent directories as needed
func Mkdir(dir string, parents bool) error {
	if Exist(dir) {
		return nil
	}

	if parents {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return os.Mkdir(dir, os.ModePerm)
}

//Exist check the given path exists
func Exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 获取指定路径的目录
func Dir(path string) string {
	return filepath.Dir(path)
}

// 获取当前绝对路径(特指运行路径，而非当前文件路径)
func GetCurrentDir() string {
	dir, _ := os.Getwd()
	return dir
}
