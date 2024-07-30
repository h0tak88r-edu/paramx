package grep

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/zomasec/logz"

	"github.com/zomasec/paramx/internal/config"
	"github.com/zomasec/paramx/pkg/types"
)

var logger = logz.DefaultLogs()

func isTypeExist(tag string, types []string) bool {
	for _, t := range types {
		if strings.EqualFold(t, tag) {
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
// - tag: A string representing the bug type to search for.
// - replaceWith: A string representing the replacement value for the found parameters.
func GrepParameters(urls []string, configs []*config.Data, tag, replaceWith string) {
	tags := []string{}

	for _, cfg := range configs {
		tags = append(tags, cfg.Tag)
	}

	if !isTypeExist(tag, tags) {
		logger.FATAL("Invalid tag , please add a valid tag like (xss, ssrf, sqli, lfi, rce, idor, ssti, redirect, isubs)")
		os.Exit(1)
	}

	for _, rawURL := range urls {
		params, fullURL := extractParameters(rawURL, replaceWith)

		for _, cfg := range configs {

			if !(cfg.Part == types.Query.String()) {
				continue
			}

			for _, param := range cfg.List {

				if strings.EqualFold(cfg.Tag, tag) {

					if _, exists := params[param]; exists {
						fmt.Println(fullURL)
					}
				}
			}
		}
	}
}

func extractParameters(rawURL, replaceWith string) (map[string]string, string) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, ""
	}

	if parsedURL.RawQuery != "" {
		query := parsedURL.Query()
		for key := range query {
			query.Set(key, replaceWith)
		}
	
		parsedURL.RawQuery = query.Encode()
	}


	params := make(map[string]string)

	for key, values := range parsedURL.Query() {	
		params[key] = values[0]
	}
	
	return params, parsedURL.String()
}

func GrepSubdomains(urls []string, configs []*config.Data) {
	for _, rawURL := range urls {

		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			continue
		}

		for _, cfg := range configs {
			if cfg.Part == types.Subdomain.String() {
				for _, subName := range cfg.List {
					if strings.Contains(parsedURL.Host, subName) {
						fmt.Println(parsedURL)
					}
				}
			}
		}
	}
}
