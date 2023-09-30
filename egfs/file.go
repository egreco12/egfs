package egfs

import "fmt"

type File struct {
	content []byte
}

func (file *File) Overwrite(c []byte) {
	file.content = c
}

func (file *File) Append(c []byte) {
	file.content = append(file.content, c...)
}

func (file *File) PrintContent() {
	fmt.Print(string(file.content))
}
