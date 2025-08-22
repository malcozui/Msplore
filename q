package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path, err := ParseCommandLineForPath()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	directoryContents, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Directory Contents:")
	RecursivelyPrintDir(path, directoryContents)
	fmt.Println()

}

func RecursivelyPrintDir(parentPath string, dirContents []os.DirEntry) error {
	for _, dirEntry := range dirContents {
		if file, err := dirEntry.Info(); err != nil {
			fmt.Println(file.Name())
			continue
		} else if dirEntry.IsDir() {
			fmt.Println(dirEntry.Name() + "/")

			newPath := parentPath + "/" + dirEntry.Name()
			subContents, err := os.ReadDir(newPath)
			if err != nil {
				return err
			}
			return RecursivelyPrintDir(newPath, subContents)
		}
	}
	return nil
}

func ParseCommandLineForPath() (string, error) {

	args := os.Args // First Arg 0 is always the executable itself. so Msplore = 0,path = 1, etc...
	if len(args) != 2 {
		return "", errors.New("You must provide only one arg, the path")
	}
	// TODO: Error handling if it's not a valid path

	path, err := GetPathCleaned(args[1])

	if err != nil {
		return "", err
	}

	return path, nil
}

func GetPathCleaned(path string) (string, error) {
	cleanedPath := filepath.Clean(path)
	fileInfo, err := os.Stat(cleanedPath)

	if path == "." {
		// . is shorthand for current dir from the CLI
		return os.Getwd()
	} else if err == nil && fileInfo.IsDir() {
		return cleanedPath, nil
	} else {
		return "", errors.New("Invalid File Path Specified")
	}
}
