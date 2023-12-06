package main

import (
	"bufio"
	"fmt"
	"os"
)

func Read(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, err
	}

	return lines, nil

}

func FindIndexFunc(slice []string, f func(string) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1 // Return -1 if no element meets the condition
}
