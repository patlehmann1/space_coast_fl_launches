package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL    = "https://ll.thespacedevs.com/2.2.0/launch/upcoming/"
	// Location ID 12 = Cape Canaveral (covers all KSC + CCSFS pads)
	locationID = "12"
	userAgent  = "fl-launches/1.0 (https://github.com/plehmann/fl-launches)"
)

type Launch struct {
	Name   string    `json:"name"`
	Net    time.Time `json:"net"`
	Status struct {
		Name string `json:"name"`
	} `json:"status"`
	Rocket struct {
		Configuration struct {
			FullName string `json:"full_name"`
		} `json:"configuration"`
	} `json:"rocket"`
	Pad struct {
		Name     string `json:"name"`
		Location struct {
			Name string `json:"name"`
		} `json:"location"`
	} `json:"pad"`
	Mission *struct {
		Description string `json:"description"`
	} `json:"mission"`
}

type response struct {
	Results []Launch `json:"results"`
}

func FetchUpcoming(n int) ([]Launch, error) {
	url := fmt.Sprintf("%s?limit=%d&location__ids=%s&ordering=net", baseURL, n, locationID)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to reach Launch Library API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("rate limit reached — try again in a few minutes")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Results, nil
}
