package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Step 1: Read the IPs from the first file
	ips, err := readLines("ips_to_check.txt")
	if err != nil {
		fmt.Println("Error reading IPs:", err)
		return
	}

	// Step 2: Read the contents of the second file
	fileContent, err := readFile("ips.txt")
	if err != nil {
		fmt.Println("Error reading file with strings:", err)
		return
	}

	// Step 3: Check each IP and print the result
	for _, ip := range ips {
		if strings.Contains(fileContent, ip) {
			fmt.Printf("%s ip exists\n", ip)
		} else {
			fmt.Printf("%s ip does not exist\n", ip)
		}
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
