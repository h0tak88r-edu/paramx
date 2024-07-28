package main

import (
	"flag"
	"os"

	"github.com/zomasec/logz"

	"github.com/zomasec/paramx/internal/runner"
	"github.com/zomasec/paramx/pkg/utils"
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

/*
https://chatgpt.com/c/68968ecd-34cf-4f0c-82a8-09e50e18e535


https://zomasec.app.spacelift.dev/stacks?sort=starred&sortDirection=DESC

https://hackerone.com/bugs?subject=user&report_id=0&view=open&substates%5B%5D=new&substates%5B%5D=needs-more-info&substates%5B%5D=pending-program-review&substates%5B%5D=triaged&substates%5B%5D=pre-submission&substates%5B%5D=retesting&reported_to_team=&text_query=&program_states%5B%5D=2&program_states%5B%5D=3&program_states%5B%5D=4&program_states%5B%5D=5&sort_type=latest_activity&sort_direction=descending&limit=25&page=1

https://hackerone.com/search?q=zomasec

*/
