package utils

import (
	"os"
)

func Exists(path string,t string) (bool) {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	if t != ""{
		typeName := "file"
		if info.IsDir() {
			typeName = "dir"
		}
		if typeName == t{
			return true
		}
	}

	return false
}

//创建文件夹
func CheckOrCreateDir(filePath string) error {
	if  !Exists(filePath,"dir"){
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}
