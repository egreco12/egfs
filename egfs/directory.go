package egfs

type Directory struct {
	Directories []*Directory
	Files       []*File
	Name        string
}
