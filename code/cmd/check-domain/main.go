package main

import (
	"fmt"
	"os"
	"strings"

	"start-feishubot/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: check-webpage <domain>")
		fmt.Println("Example: check-webpage winbests.net")
		fmt.Println("         check-webpage https://winbests.net")
		os.Exit(1)
	}

	target := os.Args[1]
	fmt.Printf("Checking webpage: %s\n", target)
	fmt.Println("=====================================")

	var isOnline bool
	var message string
	var statusCode int

	// Check if it's a full URL or just a domain
	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
		isOnline, message, statusCode = utils.CheckURLDisplay(target)
	} else {
		isOnline, message, statusCode = utils.CheckWebpageDisplay(target)
	}
	
	fmt.Printf("Status: ")
	if isOnline {
		fmt.Printf("✓ ONLINE (HTTP %d)\n", statusCode)
	} else {
		fmt.Println("✗ OFFLINE/UNREACHABLE")
	}
	fmt.Printf("Details: %s\n", message)
	
	if !isOnline {
		os.Exit(1)
	}
}
