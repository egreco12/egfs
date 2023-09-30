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
	root := egfs.Entity{Entities: make(map[string]*egfs.Entity), File: nil, Name: ""}
	egfs := egfs.EGFileSystem{Cwd: &root, Root: &root, CwdPath: ""}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s=> ", egfs.CwdPath)
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
