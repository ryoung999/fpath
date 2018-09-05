package main

import (
	"fmt"
	"os"
"strings"
"github.com/atotto/clipboard"
	"path"
)

func abspath(fpath string) (apath string) {
	cwd, err := os.Getwd()
	if err != nil {
		panic("cannot get current working directory")
	}

	if path.IsAbs(fpath) {
		apath = fpath
	} else {
		apath = path.Join(cwd, fpath)
	}
	return
}

func main() {
	paths := os.Args[1:]
	var fpaths []string

	for _, f := range paths {
		fpath := abspath(f)
		fmt.Println(fpath)
		fpaths = append(fpaths, fpath)
	}

	err := clipboard.WriteAll(strings.Join(fpaths, "\n"))

	if err != nil {
		fmt.Println("cannot copy to clipboard")
	}
}
