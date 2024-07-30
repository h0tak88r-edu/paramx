package runner

import (
	"fmt"
	"os"

	"github.com/zomasec/logz"

	"github.com/zomasec/paramx/internal/config"
	"github.com/zomasec/paramx/pkg/grep"
	"github.com/zomasec/paramx/pkg/utils"
)

var logger = logz.DefaultLogs()

// Run executes the main logic of the program.
// It downloads templates, loads configurations, and performs parameter replacement.
func Run(opts *Options) {

	if err := config.DownloadTempletes(); err != nil {
		logger.ERROR("Failed to clone repository: %s\n", err.Error())
		os.Exit(1)
	}

	if opts.TempletesPath == "" {
		opts.TempletesPath = config.TempletesPath
	}

	configs, err := config.LoadConfig(opts.TempletesPath)
	if err != nil {
		panic(err)
	}

	if opts.CustomTemplete != "" {
		date, err := config.ReadCustomTemplete(opts.CustomTemplete)
		if err != nil {
			logger.ERROR("Error reading custom templete the syntax is invalid : %s\n", err.Error())
			os.Exit(1)
		}
		configs = append(configs, date)

	}

	switch opts.Tag {
	case "isubs":
		grep.GrepSubdomains(opts.URLs, configs)
	default:
		result := utils.RemoveDuplicates(grep.GrepParameters(opts.URLs, configs, opts.Tag, opts.ReplaceWith))
	
		for _, r := range result {
			fmt.Fprintln(os.Stdout, r)
		}	
	
	}

}
