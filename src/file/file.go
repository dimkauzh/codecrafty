package file

import "os"

type File struct {
	Name string
	file *os.File
}

func newFile(name string) File {

	file_name := name + ".codecrafty"

	file, err := os.Create(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return File{name, file}
}

func writeFile(file File, text string) {
	file.file.WriteString(text)
}
