package egfs

type Node struct {
	Nodes map[string]*Node
	Name  string
	File  *File
}
