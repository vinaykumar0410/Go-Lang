package main

import (
    "fmt"
    "net/url"
)

func main() {
    
    rawURL := "https://example.com/path?query=value&param1=123#fragment"

    // Parse the URL
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        panic(err)
    }

    // Access and print URL components
    fmt.Println("Scheme:", parsedURL.Scheme) // "http"
    fmt.Println("Host:", parsedURL.Host) // "example.com"
    fmt.Println("Path:", parsedURL.Path) // "/path"
    fmt.Println("Query:", parsedURL.RawQuery) // "query=value&param1=123"
    fmt.Println("Fragment:", parsedURL.Fragment) // "fragment"

    // Modify the URL
    parsedURL.RawQuery = "modified_query=new_value"
    parsedURL.Fragment = ""

    // Reconstruct the modified URL
    modifiedURL := parsedURL.String() 
    fmt.Println("Modified URL:", modifiedURL) // "https://example.com/path?modified_query=new_value#fragment"

    // Validate the URL
    if parsedURL.IsAbs() {
        fmt.Println("URL is absolute")
    } else {
        fmt.Println("URL is relative")
    }

	// URL is absolute
	
}