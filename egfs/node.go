package egfs

type Node struct {
	Nodes       []*Node
	Name        string
	IsDirectory bool
	File        *File
}
