package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SpaceData struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    struct {
		Name           string `json:"name"`
		DisplayName    string `json:"displayName"`
		CurrentCycle   int    `json:"currentCycle"`
		CycleStartDate string `json:"cycleStartDate"`
		CurrentEvent   struct {
			Title string `json:"title"`
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"currentEvent"`
		SnapshotSpace     string `json:"snapshotSpace"`
		JuiceboxProjectId string `json:"juiceboxProjectId"`
	} `json:"data"`
}

// Get information about a nanceSpace from the nance api
func nanceSpace(space string) (*SpaceData, error) {
	resp, err := http.Get(nanceEndpoint + "/" + space)
	if err != nil {
		return nil, fmt.Errorf("error fetching from nance api: %v", err)
	}

	defer resp.Body.Close()
	var r SpaceData
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("error decoding nance api response: %v", err)
	}

	if r.Error != "" {
		return nil, fmt.Errorf("nance api returned an error: %v", r.Error)
	}

	return &r, nil
}
