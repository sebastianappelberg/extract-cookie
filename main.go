package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sebastianappelberg/extract-cookie/extract"
	"log"
	"time"
)

func main() {
	browserFlag := flag.String("browser", "firefox", "Browser to extract cookies from: chrome, chromium, edge, firefox.")
	browserShort := flag.String("b", "firefox", "Browser to extract cookies from (short).")

	formatFlag := flag.String("format", "json", "Output format: json or text.")
	formatShort := flag.String("f", "json", "Output format (short).")

	flag.Parse()

	// Prefer short flag if given
	selectedBrowser := firstNonEmpty(*browserShort, *browserFlag)
	selectedOutput := firstNonEmpty(*formatShort, *formatFlag)

	if selectedBrowser == "" {
		log.Fatal("You must specify a browser with -b or --browser")
	}
	if selectedOutput == "" {
		log.Fatal("You must specify an output format with -f or --format")
	}

	result, err := extract.Cookies(selectedBrowser)

	if err != nil {
		log.Fatalf("Error extracting cookies: %v", err)
	}

	switch selectedOutput {
	case "json":
		err = outputJSON(result)
	case "text":
		err = outputText(result)
	default:
		log.Fatalf("Unsupported output format: %s", selectedOutput)
	}

	if err != nil {
		log.Fatalf("Error writing output: %v", err)
	}
}

func outputJSON(res extract.Result) error {
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func outputText(res extract.Result) error {
	for _, c := range res.Cookies {
		fmt.Printf("Name: %s\nValue: %s\nDomain: %s\nPath: %s\nExpires: %s\nSecure: %t\nHttpOnly: %t\n\n",
			c.Name, c.Value, c.Domain, c.Path, time.Unix(c.Expires, 0).Format("2006-01-02 15:04:05"), c.Secure, c.HttpOnly)
	}
	return nil
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
