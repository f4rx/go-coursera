package main

import (
	"io"
	"io/ioutil"
	"os"
	m_path "path"
	"strconv"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	PrintFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, PrintFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTreeRecursion(out, path, printFiles, "")

}

func dirTreeRecursion(out io.Writer, path string, printFiles bool, ident string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if !printFiles {
		var dirs []os.FileInfo
		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, file)
			}
		}
		files = dirs
	}

	l := len(files) - 1
	for i, file := range files {
		prefixMark := `├───`
		if l == i {
			prefixMark = `└───`
		} else {
		}

		out.Write([]byte(ident))
		out.Write([]byte(prefixMark))
		out.Write([]byte(file.Name()))

		if file.IsDir() {
			out.Write([]byte("\n"))

			var nextIdent string
			if l == i {
				nextIdent = ident + "\t"
			} else {
				nextIdent = ident + "│\t"
			}

			err := dirTreeRecursion(out, m_path.Join(path, file.Name()), printFiles, nextIdent)
			if err != nil {
				return err
			}

		} else if printFiles {

			fileSize := file.Size()
			var fileSizeStr string
			if fileSize == 0 {
				fileSizeStr = "empty"
			} else {
				fileSizeStr = strconv.FormatInt(fileSize, 10) + "b"
			}
			out.Write([]byte(" (" + fileSizeStr + ")"))
			out.Write([]byte("\n"))
		}
	}
	return nil
}
