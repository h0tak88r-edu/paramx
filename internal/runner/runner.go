package runner

import (
    "paramx/internal/config"
    "paramx/pkg/grep"
)

func Run(urls []string, configDir string) {
    configs, err := config.LoadConfig(configDir)
    if err != nil {
        panic(err)
    }
    grep.GrepParameters(urls, configs)
}
