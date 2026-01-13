package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// CheckWebpageDisplay checks if a webpage displays normally by attempting HTTP/HTTPS connection
func CheckWebpageDisplay(domain string) (bool, string, int) {
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Allow up to 10 redirects
			if len(via) >= 10 {
				return fmt.Errorf("too many redirects")
			}
			return nil
		},
	}

	// Try HTTPS first (most common for modern websites)
	httpsURL := fmt.Sprintf("https://%s", domain)
	resp, err := client.Get(httpsURL)
	if err == nil {
		defer resp.Body.Close()
		// Read a bit of content to verify page loads
		bodyPreview := make([]byte, 100)
		_, readErr := resp.Body.Read(bodyPreview)
		if readErr != nil && readErr != io.EOF {
			return false, fmt.Sprintf("Failed to read page content: %v", readErr), 0
		}
		
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			return true, fmt.Sprintf("Webpage displays normally via HTTPS (Status: %d %s)", resp.StatusCode, http.StatusText(resp.StatusCode)), resp.StatusCode
		}
		return false, fmt.Sprintf("Webpage returned error status via HTTPS: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)), resp.StatusCode
	}

	// Try HTTP if HTTPS fails
	httpURL := fmt.Sprintf("http://%s", domain)
	resp, err = client.Get(httpURL)
	if err == nil {
		defer resp.Body.Close()
		// Read a bit of content to verify page loads
		bodyPreview := make([]byte, 100)
		_, readErr := resp.Body.Read(bodyPreview)
		if readErr != nil && readErr != io.EOF {
			return false, fmt.Sprintf("Failed to read page content: %v", readErr), 0
		}
		
		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			return true, fmt.Sprintf("Webpage displays normally via HTTP (Status: %d %s)", resp.StatusCode, http.StatusText(resp.StatusCode)), resp.StatusCode
		}
		return false, fmt.Sprintf("Webpage returned error status via HTTP: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)), resp.StatusCode
	}

	return false, fmt.Sprintf("Cannot connect to webpage: %v", err), 0
}

// CheckURLDisplay checks if a specific URL displays normally
func CheckURLDisplay(url string) (bool, string, int) {
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return fmt.Errorf("too many redirects")
			}
			return nil
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return false, fmt.Sprintf("Cannot connect to URL: %v", err), 0
	}
	defer resp.Body.Close()

	// Read a bit of content to verify page loads
	bodyPreview := make([]byte, 100)
	_, readErr := resp.Body.Read(bodyPreview)
	if readErr != nil && readErr != io.EOF {
		return false, fmt.Sprintf("Failed to read page content: %v", readErr), 0
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true, fmt.Sprintf("Webpage displays normally (Status: %d %s)", resp.StatusCode, http.StatusText(resp.StatusCode)), resp.StatusCode
	}
	return false, fmt.Sprintf("Webpage returned error status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode)), resp.StatusCode
}
