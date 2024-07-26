package runner

import (
    "paramx/internal/config"
    "paramx/pkg/grep"
)

func Run(opts *Options) {
    configs, err := config.LoadConfig(opts.ConfigPath)
    if err != nil {
        panic(err)
    }
    grep.GrepParameters(opts.URLs, configs, opts.BugType)
}
