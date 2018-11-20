package main

import (
	"fmt"
	"io"
	"os"
	// "path/filepath"
	// "strings"
	"io/ioutil"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error{
	// fmt.Println(out, path, dirTree)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return(err)
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
	return nil
}
