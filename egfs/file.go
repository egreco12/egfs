package egfs

type File struct {
	content []byte
	Name    string
}

func (file *File) Overwrite(c []byte) {
	file.content = c
}

func (file *File) Append(c []byte) {
	file.content = append(file.content, c...)
}
