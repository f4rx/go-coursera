package main

import (
	// "fmt"
	"io"
	"os"
	// "path/filepath"
	// "strings"
	"io/ioutil"
	m_path "path"
	"strconv"
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

func dirTree(out io.Writer, path string, printFiles bool) error {
	// fmt.Println(out, path, dirTree)
	// os.Stderr.Write([]byte(path + "\n")
	return dirTreeRecursion(out, path, printFiles, 0, "")

}

func dirTreeRecursion(out io.Writer, path string, printFiles bool, recCount int, ident string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return (err)
	}
	l := len(files) - 1
	for i, file := range files {
		prefixMark := "├───"
		if l == i {
			prefixMark = "└───"
		} else {
		}
		// out.Write([]byte(strings.Repeat("        ", recCount)))
		out.Write([]byte(ident))
		out.Write([]byte(prefixMark))
		out.Write([]byte(file.Name()))


		if file.IsDir() {
					out.Write([]byte("\n"))

			// out.Write([]byte(strings.Repeat("        ", recCount)))
			var next_ident string
			if l == i {
				next_ident =  ident + "        "
			} else {
				next_ident =  ident + "|       "
			}

			err := dirTreeRecursion(out, m_path.Join(path, file.Name()), printFiles, recCount+1, next_ident)
			if err != nil {
				return (err)
			}

		} else {
			fileSize := file.Size()
			var fileSizeStr string
			if fileSize == 0 {
				fileSizeStr = "empty"
			} else {
				fileSizeStr := strconv.FormatInt(fileSize, 10)
			}

			out.Write([]byte(" (" + fileSizeStr + "b)"))
			out.Write([]byte("\n"))

		}
	}
	return nil
}
