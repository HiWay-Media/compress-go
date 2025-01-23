package compress

import "strings"

// NormalizeURL removes multiple slashes from a URL and ensures it is correctly formatted.
func NormalizeURL(url string) string {
	// Replace multiple slashes with a single slash
	normalizedURL := strings.ReplaceAll(url, "//", "/")

	// Ensure the URL starts with "http://" or "https://"
	if strings.HasPrefix(url, "http://") {
		normalizedURL = "http://" + strings.TrimPrefix(normalizedURL, "http:/")
	} else if strings.HasPrefix(url, "https://") {
		normalizedURL = "https://" + strings.TrimPrefix(normalizedURL, "https:/")
	}

	return normalizedURL
}
