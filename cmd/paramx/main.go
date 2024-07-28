package main

import (
	"flag"
	"os"

	"github.com/zomasec/logz"

	"paramx/internal/runner"
	"paramx/pkg/utils"
)

var opts = &runner.Options{}
var logger = logz.DefaultLogs()

func init() {
	flag.StringVar(&opts.TempletesPath, "tp", "", "Directory where YAML configuration files are located.")
	flag.StringVar(&opts.FileInput, "l", "", "Path to a file containing URLs (one per line).")
	flag.StringVar(&opts.Tag, "tag", "xss", "The type of bug to extract the URLs based on it (default: \"xss\"). Supported values: xss, sqli, lfi, rce, idor, ssrf, ssti, redirect.")
	flag.StringVar(&opts.ReplaceWith, "rw", "", "Replace the parameter value with a custom value.")
	flag.StringVar(&opts.CustomTemplete, "t", "", "Path to a custom template.")

}

func main() {
	flag.Parse()

	
	// Read URLs from file if provided
	if opts.FileInput != "" {
		urls, err := utils.ReadURLsFromFile(opts.FileInput)
		if err != nil {
			logger.FATAL("Failed to read URLs from file: %v", err)
			os.Exit(1)
		}
		opts.URLs = urls
	} else {
		// If no file is provided, read URLs from stdin
		urls, err := utils.ReadURLsFromStdin()
		if err != nil {
			logger.FATAL("Failed to read URLs from stdin or no provided input: %v", err)
			os.Exit(1)
		}
		opts.URLs = urls
	}

	runner.Run(opts)

}
