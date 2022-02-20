package test1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Lines should be wrapped at 100 characters

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
		violations := 0
		lines := []int{}
		scanner := bufio.NewScanner(file)
		lineCount := 1

		// Scans open file for lines with over 100 characters
		for scanner.Scan() {
			if utf8.RuneCountInString(scanner.Text()) >= 100 {
				violations++
				lines = append(lines, lineCount)
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
			fmt.Println("Test 1 failed " + strconv.Itoa(totalViolations) + " times")
		} else {
			fmt.Println("Test 1 Failed " + strconv.Itoa(totalViolations) + " time")
		}
	} else if !detail {
		// if test did not fail, tells users it passed
		fmt.Println("Test 1 passed all files")
	}
	if detail {
		fmt.Println()
	}
}
