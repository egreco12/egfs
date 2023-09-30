package egfs

import (
	"fmt"
	"slices"
	"strings"
)

const (
	READ  = "read"
	WRITE = "write"
)

type File struct {
	// Contains content of file
	content []byte
	// Contains Permissions to file.  Map key is role, value is set of Permissions on that role (read or write)
	Permissions map[string][]string
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

// Checks permission on a file.  Currently only read and write are allowd;
// if isRead is false, we assume we are checking write perms, etc.
func (file *File) CheckPermission(user User, isRead bool) bool {
	perms, exists := file.Permissions[user.Role]
	if !exists {
		return false
	}

	if isRead {
		return slices.Contains(perms, READ)
	} else {
		return slices.Contains(perms, WRITE)
	}
}

// Sets permission for a given role on the file.  Permissions can be read or write.
func (file *File) SetRolePermission(role string, perm string) {
	if perm != READ && perm != WRITE {
		fmt.Printf("Error: Invalid permission provided: %s", perm)
		return
	}

	file.Permissions[role] = append(file.Permissions[role], perm)
}

// Prints the Permissions for a given role.
func (file *File) PrintPermissions(role string) {
	length := len(file.Permissions)
	builder := strings.Builder{}
	builder.WriteString("[")
	index := 0
	for _, perm := range file.Permissions[role] {
		builder.WriteString(perm)
		if index < length-1 {
			builder.WriteString(",")
		}

		index++
	}
	builder.WriteString("]\n")
	fmt.Print(builder.String())
}
