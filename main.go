package main

import (
	"errors"
	"fmt"
	// "io/fs"
	"os"
	"path/filepath"
)

func main() {

	/*
	 *		TODO: import the go package fyne for UI as it seems to be the easiest fit and is almost entirely go.
	 *	   Once that is in, then convert this code from a bad implementation of the 'ls' command,
	 *	   to a file explorer without navigation as a first step. Then move down the pipeline of adding navigation,
	 *	   adding adding opening, closing, deleting, and creating files. Also look at adding a terminal feature
	 *	   that stays synced with the currrent explored to directory (Two-way sync). Then add a file preview simular
	 *	   to that of Telescope in NeoVim so that the file explorer is useful and functional. features like drag-n-drop
	 *	   can come later, since the idea is that this file explorer would work nicely with vim motions and that
	 *	   keyboard-centric design. Another Key feature I'd like to implement is bookmarks like windows file explorer,
	 *	   but simular to vim with Harpoon, you can keybind to these specifc folders. I'd also like some sort of
	 *	   variables for file paths like an insertable environment variable, expect they're specific to the app,
	 *	   and it doesn't clutter the environment variables with additonal ones. The idea would be something like:
	 *	   *In some settings menu i'm not sure yet*: %variablename% = PathA
	 *	   C:/Users/user/%variablename% evals to C:/Users/user/PathA, useful for if you have multiple bookmarks
	 *	   to the same named folder in many different places, and you want to update all those references at once.
	 *	   this is useful for programmers in companies where one project may have sections split across one massive proj
	 *	   and they want to neatly organize all their folders, whilst being able to switch all their folders to someone
	 *	   elses at any time quickly and easiliy.
	 */

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
	RecursivelyPrintDir(directoryContents)
	fmt.Println()

}

func RecursivelyPrintDir(dirContents []os.DirEntry) error {

	for _, dirEntry := range dirContents {
		appendRune := ""
		if dirEntry.IsDir() {
			appendRune = "/"
		}
		fmt.Println(dirEntry.Name() + appendRune)
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
