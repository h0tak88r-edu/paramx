package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// var homeDir, _ = os.UserHomeDir()
// var ConfigPath = filepath.Join(homeDir, ".config", "paramx", "config")

type Bug int8

const (
    XSS Bug = iota
    SSRF
    IDOR
    SQLI
    SSTI
    OpenRedirect
    LFI
    RCE
)

type Data struct {
    BugType    string   `yaml:"bug_type"`
    Parameters []string `yaml:"parameters"`
}

func LoadConfig(configDir string) ([]*Data, error) {
    var configs []*Data

    files, err := os.ReadDir(configDir)
    if err != nil {
        return nil, err
    }

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".yaml" {
            configData, err := os.ReadFile(filepath.Join(configDir, file.Name()))
            if err != nil {
                log.Printf("error reading file: %v", err)
                continue
            }

            var data Data
            if err := yaml.Unmarshal(configData, &data); err != nil {
                log.Printf("error unmarshaling file: %v", err)
                continue
            }

            configs = append(configs, &data)
        }
    }

    return configs, nil
}
