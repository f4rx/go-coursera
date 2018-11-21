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
	return dirTreeRecursion(out, path, "", printFiles)

}

func dirTreeRecursion(out io.Writer, path, ident string, printFiles bool) (err error) {
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
		}

		_, err = out.Write([]byte(ident))
		if err != nil {
			return err
		}
		_, err = out.Write([]byte(prefixMark))
		if err != nil {
			return err
		}
		_, err = out.Write([]byte(file.Name()))
		if err != nil {
			return err
		}

		if file.IsDir() {
			_, err = out.Write([]byte("\n"))
			if err != nil {
				return err
			}

			nextIdent := ident + "│\t"
			if l == i {
				nextIdent = ident + "\t"
			}

			err := dirTreeRecursion(out, m_path.Join(path, file.Name()), nextIdent, printFiles)
			if err != nil {
				return err
			}

		} else if printFiles {

			fileSize := file.Size()
			fileSizeStr := "empty"
			if fileSize != 0 {
				fileSizeStr = strconv.FormatInt(fileSize, 10) + "b"
			}
			_, err = out.Write([]byte(" (" + fileSizeStr + ")"))
			if err != nil {
				return err
			}
			_, err = out.Write([]byte("\n"))
			if err != nil {
				return err
			}
		}
	}
	return
}
