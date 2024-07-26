package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(filePath string) ([]byte, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        log.Printf("error reading file: %v", err)
        return nil, err
    }
    return data, nil
}

// readURLsFromFile reads URLs from a file, one URL per line
func ReadURLsFromFile(filePath string) ([]string, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    return strings.Split(strings.TrimSpace(string(data)), "\n"), nil
}

// readURLsFromStdin reads URLs from standard input
func ReadURLsFromStdin() ([]string, error) {
    var urls []string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        urls = append(urls, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return urls, nil
}
