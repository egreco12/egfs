package egfs

// Represents a User in EGFS.  For now, Users have exactly
// one role for permissions reasons.
type User struct {
	Name string
	Role string
}
