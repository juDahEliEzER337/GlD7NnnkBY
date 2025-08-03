// 代码生成时间: 2025-08-03 22:58:26
package main

import (
    "fmt"
    "io/fs"
    "log"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

// Folder represents a folder structure with name, path and subfolders.
type Folder struct {
    Name     string
    Path     string
    Subs     []Folder
    Files   []string
}

// NewFolder creates a new Folder instance.
func NewFolder(name, path string) *Folder {
    return &Folder{Name: name, Path: path}
}

// AddSub adds a subfolder to the current folder.
func (f *Folder) AddSub(subFolder *Folder) {
    f.Subs = append(f.Subs, *subFolder)
}

// AddFile adds a file to the current folder.
func (f *Folder) AddFile(file string) {
    f.Files = append(f.Files, file)
}

// Traverse recursively traverses the folder and builds the folder structure.
func (f *Folder) Traverse() {
    // Read the content of the current folder
    entries, err := os.ReadDir(f.Path)
    if err != nil {
        log.Fatalf("Error reading directory: %v", err)
    }
    for _, entry := range entries {
        fullPath := filepath.Join(f.Path, entry.Name())
        if entry.IsDir() {
            // Create a new folder and add it to the current folder's subs
            subFolder := NewFolder(entry.Name(), fullPath)
            f.AddSub(subFolder)
            subFolder.Traverse()
        } else {
            f.AddFile(entry.Name())
        }
    }
}

// Print prints the folder structure.
func (f *Folder) Print(indent int) {
    fmt.Printf("%s%s/
", strings.Repeat("  ", indent), f.Name)
    for _, sub := range f.Subs {
        sub.Print(indent + 1)
    }
    for _, file := range f.Files {
        fmt.Printf("%s  %s
", strings.Repeat("  ", indent), file)
    }
}

func main() {
    // Define the root folder
    root := NewFolder("root", ".")
    // Traverse the folder structure starting from the root
    root.Traverse()
    // Sort the files and subfolders for a consistent output
    sort.SliceStable(root.Files, func(i, j int) bool { return root.Files[i] < root.Files[j] })
    sort.SliceStable(root.Subs, func(i, j int) bool { return root.Subs[i].Name < root.Subs[j].Name })
    // Print the folder structure
    root.Print(0)
}
