package grep

import (
	"net/url"
	"strings"

	"github.com/cyinnove/logify"
	"github.com/cyinnove/paramx/internal/config"
	"github.com/cyinnove/paramx/pkg/types"
)


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
func GrepParameters(urls []string, configs []*config.Data, tag, replaceWith string) []string {
	tags := []string{}

	for _, cfg := range configs {
		tags = append(tags, cfg.Tag)
	}

	if !isTypeExist(tag, tags) {
		logify.Fatalf("Invalid tag , please add a valid tag like (xss, ssrf, sqli, lfi, rce, idor, ssti, redirect, isubs)")
	}

	result := []string{}

	for _, rawURL := range urls {
		params, fullURL := extractParameters(rawURL, replaceWith)

		for _, cfg := range configs {

			if !(cfg.Part == types.Query.String()) {
				continue
			}

			for _, param := range cfg.List {

				if strings.EqualFold(cfg.Tag, tag) {

					if _, exists := params[param]; exists {
						result = append(result, fullURL)
					}
				}
			}
		}
	}

	return result
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

func GrepSubdomains(urls []string, configs []*config.Data) []string {
	result := []string{}
	for _, sub := range urls {

		for _, cfg := range configs {
			if cfg.Part == types.Subdomain.String() {
				for _, subName := range cfg.List {
					if strings.Contains(sub, subName) {
						result = append(result, sub)
					}
				}
			}
		}
	}

	return result
}
