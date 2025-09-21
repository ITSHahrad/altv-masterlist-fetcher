package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	Servers "github.com/itshahrad/altv-fetcher/internal/configs"
	json_saver "github.com/itshahrad/altv-fetcher/internal/json"
)

var (
	API_URL = "https://api.alt-mp.com/servers"
)

func main() {
	// Create request
	req, err := http.NewRequest("GET", API_URL, nil)
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
	var servers []Servers.Servers
	if err := json.Unmarshal(body, &servers); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	performance := time.Now()
	if _, err := json_saver.JsonSaver(servers); err != nil {
		fmt.Printf("failed to save fetched masterlist!\n %v", err)
	} else {
		fmt.Printf("%d servers fetched within: %v\nservers saved to servers.json!", len(servers), time.Since(performance))
	}
}
