package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	Configs "github.com/itshahrad/altv-fetcher/pkg/configs"
	Json "github.com/itshahrad/altv-fetcher/pkg/json"
)

func main() {
	// Create time for performance
	performance := time.Now()

	// Create request
	req, err := http.NewRequest("GET", Configs.API_URL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Parse JSON
	var servers []Configs.Servers
	if err := json.Unmarshal(body, &servers); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Request to save servers as json
	if _, err := Json.JsonSaver(servers); err != nil {
		fmt.Printf("failed to save fetched masterlist!\n %v", err)
	} else {
		fmt.Printf("%d servers fetched within: %v\nservers saved to servers.json!", len(servers), time.Since(performance))
	}
}
