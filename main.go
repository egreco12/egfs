package main

import (
	"bufio"
	egfs "egfs/egfs"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Welcome to egfs.  Enter input: \n")
	root := egfs.Directory{Directories: nil, Files: nil, Name: "/"}
	egfs := egfs.EGFileSystem{Cwd: &root, Root: &root}
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}

		input = strings.TrimSuffix(input, "\n")
		egfs.ProcessInput(input)
		// Move to new line
		fmt.Print("\n")
	}
}
