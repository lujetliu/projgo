package server

import (
	"os"
	"path/filepath"
)

var rootDir string

// inferRootDir 推断出项目根目录
func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) string
	infer = func(d string) string {
		// 确保项目根目录下存在 template 目录
		if exists(d + "/template") {
			return d
		}

		return infer(filepath.Dir(d)) // 如果 d 中不存在 template, 则在其
		// 上级目录递归查找
		// TODO: Dir 函数源码
	}

	rootDir = infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err) // TODO:
}
