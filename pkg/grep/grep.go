package grep

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/zomasec/logz"

	"paramx/internal/config"
)

var logger = logz.DefaultLogs()

func isTypeExist(bugType string, types []string) bool {
  for _, t := range types {
    if strings.EqualFold(t, bugType) {
      return true
    }
  }

  return false
}

// GrepParameters searches for parameters in the given URLs based on the provided configurations and bug type.
// It replaces the found parameters with the specified replacement string.
// The function takes in the following parameters:
// - urls: A slice of strings representing the URLs to search for parameters.
// - configs: A slice of *config.Data representing the configurations to use for parameter extraction.
// - bugType: A string representing the bug type to search for.
// - replaceWith: A string representing the replacement value for the found parameters.
func GrepParameters(urls []string, configs []*config.Data, bugType, replaceWith string) {
    types := []string{}

    for _, cfg := range configs {
        types = append(types, cfg.BugType)
    }

    if ! isTypeExist(bugType, types) {
        logger.FATAL("Invalid bug type")
        os. Exit(1)
    }
	
	for _, rawURL := range urls {
        params := extractParameters(rawURL, replaceWith)
		
        for _, cfg := range configs {
			
            for _, param := range cfg.Parameters {
		
				if strings.EqualFold(cfg.BugType, bugType) {

                	if _, exists := params[param]; exists {
                   		fmt.Println(param)
                	}
				}
            }
        }
    }
}

func extractParameters(rawURL, replaceWith string) map[string]string {
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        return nil
    }
    params := make(map[string]string)
    for key, values := range parsedURL.Query() {
        if replaceWith != "" {
            params[key] = replaceWith
        } else {
            params[key] = values[0]
        }
    }
    return params
}
