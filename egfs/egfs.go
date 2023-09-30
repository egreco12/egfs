package egfs

import (
	"fmt"
	"strings"
)

type EGFileSystem struct {
	Cwd     *Node
	Root    *Node
	CwdPath string
}

func IsValidEntity(entity string) bool {
	return strings.HasPrefix(entity, "\"") && strings.HasSuffix(entity, "\"")
}

func (egfs *EGFileSystem) ProcessInput(input string) {
	command := strings.Split(input, " ")
	if len(command) < 1 {
		fmt.Print("Error: Invalid command.")
		return
	}

	switch command[0] {
	case "change":
		egfs.ChangeDirectory(command)
	case "delete":
		egfs.Delete(command)
	case "get":
		egfs.Get(command)
	case "find":
		egfs.Find(command)
	case "make":
		egfs.Make(command)
	case "move":
		egfs.Move(command)
	case "write":
		egfs.WriteToFile(command)

	default:
		fmt.Print("Error: Invalid command.  Try again!")
	}
}

// Prints the name of entities in cwd
func (egfs *EGFileSystem) PrintCwdContents() {
	length := len(egfs.Cwd.Nodes)
	builder := strings.Builder{}
	builder.WriteString("[")
	index := 0
	for _, node := range egfs.Cwd.Nodes {
		builder.WriteString(node.Name)
		if index < length-1 {
			builder.WriteString(",")
		}

		index++
	}
	builder.WriteString("]\n")
	fmt.Print(builder.String())
}

func (egfs *EGFileSystem) Get(command []string) {
	getType := command[1]
	switch getType {
	case "working":
		egfs.GetWorkingDirectory(command)
		return
	case "file":
		egfs.GetFileContents(command)
	default:
		fmt.Print("Error: Unknown get command.")
	}
}

// Prints the name or contents of cwd
func (egfs *EGFileSystem) GetWorkingDirectory(command []string) {
	if len(command) == 3 {
		if command[2] != "directory" {
			fmt.Print("Error: Unknown Get Working command.")
			return
		}

		fmt.Printf("=> /%s", egfs.CwdPath)
		return
	}

	if command[2] != "directory" && command[3] != "contents" {
		fmt.Print("Error: Uknown Get Working command.")
		return
	}

	egfs.PrintCwdContents()
}

// Prints the contents of the provided file
func (egfs *EGFileSystem) GetFileContents(command []string) {
	entity := command[2]
	if !IsValidEntity(entity) {
		fmt.Printf("Error: Invalid entity %s: must be wrapped in quotes.", entity)
		return
	}

	node := egfs.GetNode(entity)
	if node == nil {
		fmt.Printf("Error: provided filename %s does not exist", entity)
		return
	}

	if node.File == nil {
		fmt.Print("Error: provided node is not a file.")
		return
	}

	node.File.PrintContent()
}

// Creates a new entity
func (egfs *EGFileSystem) Make(command []string) {
	if len(command) < 3 {
		fmt.Print("Error: Make command input must include at least 3 strings.")
		return
	}

	entity := command[1]
	if !IsValidEntity(entity) {
		fmt.Printf("Error: Invalid entity %s: must be wrapped in quotes.", entity)
		return
	}

	if egfs.GetNode(entity) != nil {
		fmt.Printf("Error: Entity with name %s already exists.", entity)
		return
	}

	switch command[2] {
	case "directory":
		egfs.MakeDirectory(entity)
		return

	case "file":
		egfs.MakeFile(entity)
		return
	}

	fmt.Print("Error: Invalid Make command.")
}

// Creates an empty directory under cwd
func (egfs *EGFileSystem) MakeDirectory(name string) {
	node := Node{Nodes: make(map[string]*Node), Name: name}
	egfs.Cwd.Nodes[node.Name] = &node
}

// Creates an empty file under cwd
func (egfs *EGFileSystem) MakeFile(name string) {
	node := Node{Nodes: make(map[string]*Node), File: &File{}, Name: name}
	egfs.Cwd.Nodes[node.Name] = &node
}

func (egfs *EGFileSystem) ChangeDirectory(command []string) {
	if len(command) < 4 {
		fmt.Print("Error: Invalid change directory command.  Must contain 4 strings")
		return
	}

	if command[1] != "directory" && command[2] != "to" {
		fmt.Print("Error: Invalid change directory command.  The command must be of format \"change directory to \"<new_directory>\"\"")
		return
	}

	newCwdName := command[3]
	var found bool = false
	for _, dir := range egfs.Cwd.Nodes {
		if dir.Name == newCwdName {
			newCwdName := strings.ReplaceAll(newCwdName, "\"", "")
			found = true
			if egfs.CwdPath == "" {
				egfs.CwdPath = newCwdName
			} else {
				egfs.CwdPath = fmt.Sprintf("%s/%s", egfs.CwdPath, newCwdName)
			}
			egfs.Cwd = dir
		}
	}

	if !found {
		fmt.Print("Directory not found.")
	}
}

// Deletes a file under cwd
func (egfs *EGFileSystem) Delete(command []string) {
	entity := command[1]
	if !IsValidEntity(entity) {
		fmt.Printf("Error: Invalid name %s: must be wrapped in quotes.", entity)
		return
	}

	if command[2] != "directory" {
		fmt.Print("Error: invalid entity type provided.")
	}

	// Node to delete can only be in CWD's contents
	delete(egfs.Cwd.Nodes, entity)
}

// Writes to file under cwd
func (egfs *EGFileSystem) WriteToFile(command []string) {
	name := command[1]
	node := egfs.GetNode(name)
	if node == nil {
		fmt.Printf("Error: provided entity %s does not exist", name)
		return
	}

	// Re-expand command; everything after filename is content
	node.File.Append([]byte(strings.Join(command[2:], " ")))
}

// Moves entity to new location in cwd
func (egfs *EGFileSystem) Move(command []string) {
	entity := command[1]
	if !IsValidEntity(entity) {
		fmt.Printf("Error: Invalid name %s: must be wrapped in quotes.", entity)
		return
	}

	newLoc := command[2]
	if !IsValidEntity(newLoc) {
		fmt.Printf("Error: Invalid name %s: must be wrapped in quotes.", entity)
		return
	}

	node := egfs.GetNode(entity)
	delete(egfs.Cwd.Nodes, entity)
	node.Name = newLoc
	egfs.Cwd.Nodes[newLoc] = node
}

// Returns node with provided name in cwd
func (egfs *EGFileSystem) GetNode(name string) *Node {
	node, exists := egfs.Cwd.Nodes[name]
	if !exists {
		return nil
	}

	return node
}

// Finds a file in cwd
func (egfs *EGFileSystem) Find(command []string) {
	name := command[1]
	if !IsValidEntity(name) {
		fmt.Printf("Error: Invalid name %s: must be wrapped in quotes.", name)
		return
	}

	for _, node := range egfs.Cwd.Nodes {
		if node.Name == name {
			entityType := "file"
			if node.File == nil {
				entityType = "directory"
			}

			fmt.Printf("Found entity with name %s, type: %s", name, entityType)
			return
		}
	}

	fmt.Print("Entity not found.")
}
