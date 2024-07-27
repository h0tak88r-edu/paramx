package runner

import (
	"os"

	"github.com/zomasec/logz"

	"paramx/internal/config"
	"paramx/pkg/grep"
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
    grep.GrepParameters(opts.URLs, configs, opts.BugType, opts.ReplaceWith)
}
