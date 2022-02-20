package test2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Tabs should be utilized for line indention (not spaces) and only
// line feeds utilized (LF, /n) at the end of lines

func Test(list []string, detail bool) {
	if detail {
		fmt.Println()
	}
	totalViolations := 0
	for i := range list {
		// opens file from provided list
		file, err := os.Open(list[i])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Creates necessary variables for testing
		scanner := bufio.NewScanner(file)
		violations := 0
		lines := []int{}
		lineCount := 1
		previousIndent := 0

		// Scans for lines with leading spaces and tabs
		for scanner.Scan() {
			l := scanner.Text()
			// ignore if current line is empty
			if l != "" {
				leadSpace := len(l) - len(strings.TrimLeft(l, " "))
				leadIndent := len(l) - len(strings.TrimLeft(l, "	"))
				incremConds := (leadIndent-1 != previousIndent) && (leadIndent+1 != previousIndent)
				// if spaces are used instead of tabs, return an error
				if leadSpace != 0 {
					violations++
					lines = append(lines, lineCount)
				} else if (leadIndent != previousIndent) && incremConds {
					// if tabs are used but incorrectly, returns an error
					violations++
					lines = append(lines, lineCount)
				}
				previousIndent = leadIndent
			}
			lineCount++
		}

		// Prints detailed test results
		totalViolations = totalViolations + violations
		if violations > 0 && detail {
			// variables for readability
			violationsStr := strconv.Itoa(violations)
			base := filepath.Base(list[i])
			trimmedList := strings.Trim(strings.Replace(fmt.Sprint(lines), " ", ","+" ", -1), "[]")
			// failures
			if violations > 1 {
				fmt.Println("Failed " + violationsStr + " times at " + base + " lines: " + trimmedList)
			} else {
				fmt.Println("Failed " + violationsStr + " time at " + base + " line: " + trimmedList)
			}
		} else if detail {
			// if test did not fail, tells users it passed
			fmt.Println("Passed at " + filepath.Base(list[i]))
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	// Prints general test results
	if totalViolations > 0 && !detail {
		// failures
		if totalViolations > 1 {
			fmt.Println("Test 2 failed " + strconv.Itoa(totalViolations) + " times")
		} else {
			fmt.Println("Test 2 Failed " + strconv.Itoa(totalViolations) + " time")
		}
	} else if !detail {
		// if test did not fail, tells users it passed
		fmt.Println("Test 2 passed all files")
	}
	if detail {
		fmt.Println()
	}
}
