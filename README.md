# EGFS: An In-Memory File System in Go

## Installation
* Requires Go version 1.21.1.  You can download Go for your system [here](https://go.dev/doc/install).
* After installing Go, clone the repo and navigate to the repo root.  Then, run `go run main.go` to start egfs.

## Basic commands:
* **make**: Creates a new directory or file
  * Examples:
    * File: `make "new_file" file`
    * Directory: `make "new_directory" directory`
* **change**: Changes to a directory in current working directory
  * Change to parent directory: `change directory to ..`
  * Change to directory in current working directory: `change directory to "school"`
* **get**: Returns information about a provided entity
  * Examples:
    * Get current working directory (including path): `get working directory`
    * Get contents of current working directory (includes files and directories): `get working directory contents`
    * Get contents of a file: `get file "filename"`
* **delete**: Deletes an entity (could be a file or directory)
  * Example: `delete "entity_name"`
* **find**: Finds the name of an entity within the current working directory
  * Example: `find "entity_name"`
* **write**: Writes content to a provided file.  Overwrites the provided file with all content after the provided filename.
  * Example: `write "new_file" file contents`
    * Content of file in this case is `file contents`
* **move**: Moves entity to new location in current working directory.
  * Example: `move "new_entity" "newer_entity"`

## Permissions-related commands:
Permissions are ignored if none are set on a File.  However, once a permission is set on a file,
both read and write permissions are checked when access on a file is requested.
* **user**: Set or get user of egfs.  You must provide a user name and role
  * Set user: `user set user_1 role_1`
  * Get user: `user get`
* **set**: Set permissions on a file.  Currently supports read and write.
  * Example: `set read role_1 "file_1"`
* **get**: Get permissions on a file.  Must provide a role name.
  * Example: `get permissions role_1 "file_1"`

