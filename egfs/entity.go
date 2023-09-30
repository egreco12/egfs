package egfs

type Entity struct {
	Entities map[string]*Entity
	Name     string
	File     *File
	Parent   *Entity
}
