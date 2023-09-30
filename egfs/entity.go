package egfs

// Represents an Entity in EGFS.  Generally, an Entity can be a file or a directory.
// An Entity is a file if its associated File pointer is nil.
type Entity struct {
	Entities map[string]*Entity
	Name     string
	File     *File
	Parent   *Entity
}
