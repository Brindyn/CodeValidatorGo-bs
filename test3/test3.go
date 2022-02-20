package test3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Source code must be commented with comments being valuable but conservative

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
		commentNum := 0
		started := false
		violations := 0
		lines := []int{}
		scanner := bufio.NewScanner(file)
		lineCount := 1

		// Scans for comments in main body code ensuring at least 1 comment every 15 lines
		for scanner.Scan() {
			l := scanner.Text()
			// starts counting commentless lines after imports
			if l == ")" {
				started = true
			}
			if started {
				// if a comment is found, reset the count
				if strings.Contains(l, "//") {
					commentNum = 0
				}
				// if 10 lines straight with no comments, then test fails
				if commentNum > 14 {
					violations++
					lines = append(lines, lineCount)
				}
				// increases commentless line count
				commentNum++
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
			fmt.Println("Test 3 failed " + strconv.Itoa(totalViolations) + " times")
		} else {
			fmt.Println("Test 3 Failed " + strconv.Itoa(totalViolations) + " time")
		}
	} else if !detail {
		// if test did not fail, tells users it passed
		fmt.Println("Test 3 passed all files")
	}
	if detail {
		fmt.Println()
	}
}
