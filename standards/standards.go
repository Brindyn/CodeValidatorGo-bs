package standards

import (
	"fmt"
	"github.com/Brindyn/CodeValidatorGo-bs/tree/master/test1"
	"github.com/Brindyn/CodeValidatorGo-bs/tree/master/test2"
	"github.com/Brindyn/CodeValidatorGo-bs/tree/master/test3"
	"github.com/Brindyn/CodeValidatorGo-bs/tree/master/test4"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Private list of possible commands
var commands = []string{"help", "exit", "detail", "test1", "test2", "test3", "test4", "testAll"}

// List of descriptions for each command
var commandsDesc = []string{" - lists all possible commands", " - quits the program",
	" - reports the exact locations for all failed validations of all standards",
	" - lines should be wrapped at 100 characters",
	" - tabs should be utilized for line indention (not spaces)" +
		" and only line feeds utilized (LF, /n) at the end of lines",
	" - source code must be commented with comments being valuable but conservative",
	" - all files should to be UTF-8 compatible text files",
	" - reports the number of failed validations of all of the standards"}

// Finds if the given command exists and returns an error if it does not.
func FindCommand(com string) string {
	for i, n := range commands {
		if com == n {
			// parse command to be executable if one exists
			parseCom(commands[i])
			return "nil"
		}
	}
	return "unknown command"
}

// Gets all files in the root folder
func getFiles(manual bool) []string {
	// gets the root directory from val
	_, filename, _, _ := runtime.Caller(0)
	rootpath := filename
	if manual {
		rootpath = strings.TrimSuffix(filename, "coding-standards-validator-bs/standards/standards.go")
	} else {
		rootpath = strings.TrimSuffix(filename, "standards/standards.go")
	}

	// makes a list of filepaths for every file in the directory tree
	list := make([]string, 0, 10)

	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// finds go files and appends them to the list
		if filepath.Ext(path) == ".go" {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		// finds errors with the filepath walk
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}

// Parses found command to be able to execute correct function
func parseCom(com string) {
	if com == "help" {
		// gives command list
		help()
	}
	if com == "detail" {
		// gives detailed auto testing results
		fmt.Println()
		fmt.Println("Executing detailed test 1...")
		test1.Test(getFiles(false), true)
		fmt.Println("Executing detailed test 2...")
		test2.Test(getFiles(false), true)
		fmt.Println("Executing detailed test 3...")
		test3.Test(getFiles(false), true)
		fmt.Println("Executing detailed test 4...")
		test4.Test(getFiles(false), true)
		fmt.Println("done")
		fmt.Println()
	}
	if com == "test1" {
		// manual test1
		test1.Test(getFiles(true), true)
	}
	if com == "test2" {
		// manual test2
		test2.Test(getFiles(true), true)
	}
	if com == "test3" {
		// manual test3
		test3.Test(getFiles(true), true)
	}
	if com == "test4" {
		// manual test4
		test4.Test(getFiles(true), true)
	}
	if com == "testAll" {
		// auto test all tests
		fmt.Println()
		test1.Test(getFiles(false), false)
		test2.Test(getFiles(false), false)
		test3.Test(getFiles(false), false)
		test4.Test(getFiles(false), false)
		fmt.Println()
	}
}

// Prints help text for given command
func help() {
	fmt.Println()
	for i, n := range commands {
		fmt.Println("val " + n + commandsDesc[i])
	}
	fmt.Println()
}
