package test4

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

// All files should to be UTF-8 compatible text files

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

		// Scans lines for utf8 compatible characters
		for scanner.Scan() {
			b := []byte(scanner.Text())
			// if a non utf8 compatible character is found in a line, returns an error
			if !utf8.ValidString(string(b)) {
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
			fmt.Println("Test 4 failed " + strconv.Itoa(totalViolations) + " times")
		} else {
			fmt.Println("Test 4 Failed " + strconv.Itoa(totalViolations) + " time")
		}
	} else if !detail {
		// if test did not fail, tells users it passed
		fmt.Println("Test 4 passed all files")
	}
	if detail {
		fmt.Println()
	}
}
