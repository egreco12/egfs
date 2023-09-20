package egfs

import (
	"fmt"
	"strings"
)

type EGFileSystem struct {
	Cwd  *Node
	Root *Node
}

func (egfs *EGFileSystem) ProcessInput(input string) {
	command := strings.Split(input, " ")
	if len(command) < 1 {
		fmt.Print("Invalid command.")
		return
	}

	switch command[0] {
	case "make":
		egfs.Make(command)

	case "change":
		egfs.ChangeDirectory(command)

	case "get":
		egfs.Get(command)

	default:
		fmt.Print("Invalid command.  Try again!")
	}
}

func (egfs *EGFileSystem) PrintCwdContents() {
	length := len(egfs.Cwd.Nodes)
	builder := strings.Builder{}
	builder.WriteString("[")
	for index, node := range egfs.Cwd.Nodes {
		builder.WriteString(node.Name)
		if index < length-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]\n")
	fmt.Print(builder.String())
}

func (egfs *EGFileSystem) Get(command []string) {
	if len(command) < 3 {
		fmt.Print("Error: Get command must include at least 3 strings.")
		return
	}

	if command[1] == "working" {
		if len(command) == 3 {
			if command[2] != "directory" {
				fmt.Print("Error: Unknown Get Working command.")
				return
			}

			fmt.Printf("=> %s", egfs.Cwd.Name)
			return
		}

		if command[2] != "directory" && command[3] != "contents" {
			fmt.Print("Error: Uknown Get Working command.")
			return
		}

		egfs.PrintCwdContents()
	} else {
		fmt.Print("Unknown get command.")
	}
}

func (egfs *EGFileSystem) Make(command []string) {
	if len(command) < 3 {
		fmt.Print("Error: Make command input must include at least 3 strings.")
		return
	}

	name := command[1]
	if !strings.HasPrefix(name, "\"") && !strings.HasSuffix(name, "\"") {
		fmt.Printf("Invalid name %s: must be wrapped in quotes.", name)
		return
	}
	newDir := Node{Nodes: nil, File: nil, IsDirectory: true, Name: name}
	egfs.Cwd.Nodes = append(egfs.Cwd.Nodes, &newDir)
}

func (egfs *EGFileSystem) ChangeDirectory(command []string) {
	if len(command) < 4 {
		fmt.Print("Invalid change directory command.  Must contain 4 strings")
		return
	}

	if command[1] != "directory" && command[2] != "to" {
		fmt.Print("Invalid change directory command.  The command must be of format \"change directory to \"<new_directory>\"\"")
		return
	}

	newCwdName := command[3]
	for _, dir := range egfs.Cwd.Nodes {
		if dir.Name == newCwdName {
			egfs.Cwd = dir
			fmt.Printf("=> %s", dir.Name)
		}
	}
}
