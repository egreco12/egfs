package egfs

import (
	"fmt"
	"strings"
)

type EGFileSystem struct {
	Cwd *Directory
}

func (egfs *EGFileSystem) Make(name string) {
	newDir := Directory{Directories: nil, Files: nil, Name: name}
	egfs.Cwd.Directories = append(egfs.Cwd.Directories, &newDir)
}

func (egfs *EGFileSystem) ProcessInput(input string) {
	content := strings.Split(input, " ")
	if len(content) < 3 {
		fmt.Print("Error: input must include at least 3 strings.")
		return
	}
	switch content[0] {
	case "make":
		name := content[1]
		if !strings.HasPrefix(name, "\"") && !strings.HasSuffix(name, "\"") {
			fmt.Printf("Invalid name %s: must be wrapped in quotes.", name)
			return
		}

		obj := content[2]
		switch obj {
		case "directory":
			egfs.Make(name)
			fmt.Printf("Directory created.  Current content length: %d", len(egfs.Cwd.Directories))
		}
	}
}
