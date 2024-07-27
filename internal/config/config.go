package config

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/zomasec/logz"
	"gopkg.in/yaml.v3"
)


type TemplateType int8

const (
    Subdomain TemplateType = iota
    Path
    Query

)


func (t TemplateType) String() string {
    return [...]string{"subdomain", "path", "query"}[t]
}

var logger = logz.DefaultLogs()

var TempletesPath = filepath.Join(os.Getenv("HOME"), "paramx-templetes")


type Data struct {
    Tag    string   `yaml:"tag"`
    Part  string   `yaml:"part"` // subdomain, path, query
    List []string `yaml:"list"`
}

// Check config path

func DownloadTempletes() error {
    if _, err := os.Stat(TempletesPath); os.IsNotExist(err) {
        logger.INFO("Templates directory does not exist. Cloning repository...")
        cmd := exec.Command("git", "clone", "https://github.com/zomasec/paramx-templetes.git", TempletesPath)
        err := cmd.Run()
        if err != nil {
            return err
        }
        logger.INFO("Param Templetes installed successfully.")
        return nil
    }

    return nil
}

// LoadConfig loads configuration files from the specified directory and returns a slice of Data objects.
// It reads all files with the ".yaml" extension in the directory and unmarshals them into Data objects.
// Any errors encountered during file reading or unmarshaling are logged, and the corresponding files are skipped.
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
                logger.ERROR("error reading file: %v", err)
                continue
            }

            var data Data
            if err := yaml.Unmarshal(configData, &data); err != nil {
                logger.ERROR("error unmarshaling file: %v", err)
                continue
            }

            configs = append(configs, &data)
        }
    }

    return configs, nil
}


func ReadCustomTemplete(filePath string) (*Data, error) {

    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }

    defer file.Close()

    var data Data

    decoder := yaml.NewDecoder(file)

    if err := decoder.Decode(&data); err != nil {
        return nil, err
    }

    return &data, nil

}

func UpdateTempletes () error {
    cmd := exec.Command("git", "-C", TempletesPath, "pull")
    err := cmd.Run()
    if err != nil {
        return err
    }
    logger.INFO("Param Templetes updated successfully.")
    return nil
}