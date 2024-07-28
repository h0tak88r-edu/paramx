package utils

import (
	"bufio"
	"os"

	"github.com/zomasec/logz"
)

var logger = logz.DefaultLogs()

func ReadFile(filePath string) ([]byte, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        logger.ERROR("error reading file: %s", err.Error())
        return nil, err
    }
    return data, nil
}

// readURLsFromFile reads URLs from a file, one URL per line
func ReadURLsFromFile(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }

    var urls []string

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        if scanner.Text() != "" {
            urls = append(urls, scanner.Text())
        }
    }

    return urls, nil
}

// readURLsFromStdin reads URLs from standard input
// ReadURLsFromStdin reads URLs from standard input and returns them as a slice of strings.
// It reads each line from the standard input until there are no more lines, and appends each line to the URLs slice.
// If an error occurs while reading from the standard input, it returns the error.
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

