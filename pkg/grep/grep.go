package grep

import (
	"fmt"
	"net/url"

	"paramx/internal/config"
)

func GrepParameters(urls []string, configs []config.Data) {
    for _, rawURL := range urls {
        params := extractParameters(rawURL)
		
        for _, cfg := range configs {
            for _, param := range cfg.Parameters {
				
                if _, exists := params[param]; exists {
                    fmt.Println(param)
                }
            }
        }
    }
}

func extractParameters(rawURL string) map[string]string {
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        return nil
    }
    params := make(map[string]string)
    for key, values := range parsedURL.Query() {
        params[key] = values[0]
    }
    return params
}
