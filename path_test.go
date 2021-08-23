package main

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"testing"
)

func GetFile(p string) (string,error) {
	f, err := ioutil.ReadFile(p)
	if err != nil {
		return "",err
	}
	return string(f),nil
}

func TestGetPwd(t *testing.T) {
	workPath,_:=os.Getwd()
	t.Log(workPath)
	modPath:= path.Join(workPath,"go.mod")
	fileStr,err := GetFile(modPath)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fileStr)
	re := regexp.MustCompile(`^module (.*)\n`)
	params := re.FindStringSubmatch(fileStr)
	t.Log(params)
	for _,param :=range params {
		t.Log(param)
	}
}