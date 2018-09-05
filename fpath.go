package main

import (
	"fmt"
	"os"
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

	for _, fpath := range paths {
		fmt.Println(abspath(fpath))
	}
}
