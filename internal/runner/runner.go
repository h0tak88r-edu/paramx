package runner

import (
	"fmt"
	"os"

	"github.com/cyinnove/logify"

	"github.com/cyinnove/paramx/internal/config"
	"github.com/cyinnove/paramx/pkg/grep"
	"github.com/cyinnove/paramx/pkg/utils"
)

// Run executes the main logic of the program.
// It downloads templates, loads configurations, and performs parameter replacement.
func Run(opts *Options) {

	if err := config.DownloadTempletes(); err != nil {
		logify.Errorf("Failed to clone repository: %s\n", err.Error())
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
			logify.Errorf("Error reading custom templete the syntax is invalid : %s\n", err.Error())
			os.Exit(1)
		}
		configs = append(configs, date)

	}


	switch opts.Tag {
	case "isubs":
		logify.Infof("Starting getting intersting subdomains from %d subdomains", len(opts.URLs), opts.Tag)

		result := utils.RemoveDuplicates(grep.GrepSubdomains(opts.URLs, configs))
		for _, r := range result {
			fmt.Fprintln(os.Stdout, r)
		}
		
		logify.Infof("Found %d interesting subdomains", len(result))

	default:
		logify.Infof("Starting getting parameters from %d urls for tag %s", len(opts.URLs), opts.Tag)

		result := utils.RemoveDuplicates(grep.GrepParameters(opts.URLs, configs, opts.Tag, opts.ReplaceWith))

		for _, r := range result {
			fmt.Fprintln(os.Stdout, r)
		}

		logify.Infof("Found %d parameter with tag %s", len(result), opts.Tag)

	}

}
