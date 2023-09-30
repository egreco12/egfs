package egfs

import (
	"fmt"
	"strings"
)

type EGFileSystem struct {
	Cwd     *Entity
	Root    *Entity
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
	length := len(egfs.Cwd.Entities)
	builder := strings.Builder{}
	builder.WriteString("[")
	index := 0
	for _, entity := range egfs.Cwd.Entities {
		builder.WriteString(entity.Name)
		if index < length-1 {
			builder.WriteString(",")
		}

		index++
	}
	builder.WriteString("]\n")
	fmt.Print(builder.String())
}

// Prints contents of provided entity in cwd
func (egfs *EGFileSystem) Get(command []string) {
	getType := command[1]
	switch getType {
	case "working":
		egfs.GetWorkingDirectory(command)
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

		fmt.Printf("/%s", egfs.CwdPath)
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
	name := command[2]
	if !IsValidEntity(name) {
		fmt.Printf("Error: Invalid entity %s: must be wrapped in quotes.", name)
		return
	}

	entity := egfs.GetEntity(name)
	if entity == nil {
		fmt.Printf("Error: provided entity %s does not exist", entity)
		return
	}

	if entity.File == nil {
		fmt.Print("Error: provided entity is not a file.")
		return
	}

	entity.File.PrintContent()
}

// Creates a new entity
func (egfs *EGFileSystem) Make(command []string) {
	if len(command) < 3 {
		fmt.Print("Error: Make command input must include at least 3 strings.")
		return
	}

	name := command[1]
	if !IsValidEntity(name) {
		fmt.Printf("Error: Invalid entity %s: must be wrapped in quotes.", name)
		return
	}

	if egfs.GetEntity(name) != nil {
		fmt.Printf("Error: Entity with name %s already exists.", name)
		return
	}

	var entity *Entity
	switch command[2] {
	case "directory":
		entity = egfs.MakeDirectory(name)

	case "file":
		entity = egfs.MakeFile(name)

	default:
		fmt.Print("Error: Invalid Make command.")
		return
	}

	if entity != nil {
		entity.Parent = egfs.Cwd
	}
}

// Creates an empty directory under cwd
func (egfs *EGFileSystem) MakeDirectory(name string) *Entity {
	entity := Entity{Entities: make(map[string]*Entity), Name: name}
	egfs.Cwd.Entities[entity.Name] = &entity
	fmt.Printf("Created directory %s.", entity.Name)
	return &entity
}

// Creates an empty file under cwd
func (egfs *EGFileSystem) MakeFile(name string) *Entity {
	entity := Entity{Entities: make(map[string]*Entity), File: &File{}, Name: name}
	egfs.Cwd.Entities[entity.Name] = &entity
	fmt.Printf("Created file %s.", entity.Name)
	return &entity
}

// Changes directory.  Can be to parent directory or any entitys.
func (egfs *EGFileSystem) ChangeDirectory(command []string) {
	if len(command) < 4 {
		fmt.Print("Error: Invalid change directory command.  Must contain 4 arguments")
		return
	}

	if command[1] != "directory" && command[2] != "to" {
		fmt.Print("Error: Invalid change directory command.  The command must be of format \"change directory to \"<new_directory>\"\"")
		return
	}

	newCwdName := command[3]
	// Implies change directory to parent
	if newCwdName == ".." {
		if egfs.Cwd.Parent == nil {
			fmt.Print("Error: Already at root, cant change to parent directory.")
			return
		}

		egfs.Cwd = egfs.Cwd.Parent
		// HACK: There's probably a cleaner way to do this, but if parent's parent is root, just overwrite CwdPath to ""
		if egfs.Cwd.Parent == nil {
			egfs.CwdPath = ""
		} else {
			egfs.CwdPath = strings.Join(strings.Split(egfs.CwdPath, "/")[:1], "/")
		}
	} else {
		var found bool = false
		for _, entity := range egfs.Cwd.Entities {
			if entity.Name == newCwdName {
				if entity.File != nil {
					fmt.Printf("Error: Provided entity %s is a file.  Provide a directory to change to.", entity.Name)
					return
				}

				newCwdName := strings.ReplaceAll(newCwdName, "\"", "")
				found = true
				if egfs.CwdPath == "" {
					egfs.CwdPath = newCwdName
				} else {
					egfs.CwdPath = fmt.Sprintf("%s/%s", egfs.CwdPath, newCwdName)
				}
				egfs.Cwd = entity
			}
		}

		if !found {
			fmt.Print("Directory not found.")
			return
		}
	}

	egfs.PrintCwdContents()
}

// Deletes an entity under cwd
func (egfs *EGFileSystem) Delete(command []string) {
	entity := command[1]
	if !IsValidEntity(entity) {
		fmt.Printf("Error: Invalid name %s: must be wrapped in quotes.", entity)
		return
	}

	// Entity to delete can only be in CWD's contents
	delete(egfs.Cwd.Entities, entity)
	fmt.Printf("Deleted entity %s.", entity)
}

// Writes to file under cwd
func (egfs *EGFileSystem) WriteToFile(command []string) {
	name := command[1]
	entity := egfs.GetEntity(name)
	if entity == nil {
		fmt.Printf("Error: provided entity %s does not exist", name)
		return
	}

	// Re-expand command; everything after filename is content
	entity.File.Append([]byte(strings.Join(command[2:], " ")))
	fmt.Printf("Finished writing to file %s.", name)
}

// Moves entity to new location in cwd
func (egfs *EGFileSystem) Move(command []string) {
	name := command[1]
	if !IsValidEntity(name) {
		fmt.Printf("Error: Invalid entity %s: must be wrapped in quotes.", name)
		return
	}

	newLoc := command[2]
	if !IsValidEntity(newLoc) {
		fmt.Printf("Error: Invalid entity %s: must be wrapped in quotes.", newLoc)
		return
	}

	entity := egfs.GetEntity(name)
	delete(egfs.Cwd.Entities, name)
	entity.Name = newLoc
	egfs.Cwd.Entities[newLoc] = entity
	fmt.Printf("Moved to new entity %s.", newLoc)
}

// Returns entity with provided name in cwd
func (egfs *EGFileSystem) GetEntity(name string) *Entity {
	entity, exists := egfs.Cwd.Entities[name]
	if !exists {
		return nil
	}

	return entity
}

// Finds a file in cwd
func (egfs *EGFileSystem) Find(command []string) {
	name := command[1]
	if !IsValidEntity(name) {
		fmt.Printf("Error: Invalid name %s: must be wrapped in quotes.", name)
		return
	}

	for _, entity := range egfs.Cwd.Entities {
		if entity.Name == name {
			entityType := "file"
			if entity.File == nil {
				entityType = "directory"
			}

			fmt.Printf("Found entity with name %s, type: %s", name, entityType)
			return
		}
	}

	fmt.Print("Entity not found.")
}
