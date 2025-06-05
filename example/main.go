package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sambatechno/nclib"
)

func main() {
	config := loadConfig()

	client := nclib.NewClient(nclib.US, nclib.WithDebug(true))
	timestamp := time.Now().Format(time.RFC3339)
	payload := nclib.AddActivityPayload{
		ActivityName:   "test_activity",
		AssetId:        config["asset_id"],
		Timestamp:      timestamp,
		Identity:       "d27b1360-191f-11f0-8556-4201ac16b00a",
		ActivitySource: "app",
		ActivityParams: map[string]any{
			"timestamp": timestamp,
			"param2":    "value2",
		},
	}

	if err := client.WithApiKey(config["api_key"]).AddActivity(context.Background(), payload); err != nil {
		fmt.Println("Error adding activity:", err)
	}
}

func loadConfig() map[string]string {
	file, err := os.ReadFile("example/config.json")
	if err != nil {
		panic(err)
	}
	var config map[string]string
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	return config
}
