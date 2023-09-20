package egfs

type Node struct {
	Nodes       map[string]*Node
	Name        string
	IsDirectory bool
	File        *File
}
