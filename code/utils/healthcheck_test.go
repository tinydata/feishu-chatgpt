package utils

import (
	"strings"
	"testing"
)

func TestCheckWebpageDisplay(t *testing.T) {
	tests := []struct {
		name           string
		domain         string
		expectOnline   bool
		expectErrorMsg string
	}{
		{
			name:           "Non-existent domain",
			domain:         "this-domain-definitely-does-not-exist-12345.com",
			expectOnline:   false,
			expectErrorMsg: "Cannot connect to webpage",
		},
		{
			name:           "Winbests.net domain",
			domain:         "winbests.net",
			expectOnline:   false,
			expectErrorMsg: "Cannot connect to webpage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isOnline, message, _ := CheckWebpageDisplay(tt.domain)
			
			if isOnline != tt.expectOnline {
				t.Errorf("CheckWebpageDisplay(%s) online status = %v, want %v", tt.domain, isOnline, tt.expectOnline)
			}
			
			if tt.expectErrorMsg != "" && !strings.Contains(message, tt.expectErrorMsg) {
				t.Errorf("CheckWebpageDisplay(%s) message = %v, want to contain %v", tt.domain, message, tt.expectErrorMsg)
			}
		})
	}
}

func TestCheckURLDisplay(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectOnline   bool
		expectErrorMsg string
	}{
		{
			name:           "Invalid URL",
			url:            "http://this-domain-definitely-does-not-exist-12345.com",
			expectOnline:   false,
			expectErrorMsg: "Cannot connect to URL",
		},
		{
			name:           "Malformed URL",
			url:            "not-a-valid-url",
			expectOnline:   false,
			expectErrorMsg: "Cannot connect to URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isOnline, message, _ := CheckURLDisplay(tt.url)
			
			if isOnline != tt.expectOnline {
				t.Errorf("CheckURLDisplay(%s) online status = %v, want %v", tt.url, isOnline, tt.expectOnline)
			}
			
			if tt.expectErrorMsg != "" && !strings.Contains(message, tt.expectErrorMsg) {
				t.Errorf("CheckURLDisplay(%s) message = %v, want to contain %v", tt.url, message, tt.expectErrorMsg)
			}
		})
	}
}
