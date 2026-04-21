package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var rawJSON = []byte(
	`{
  		"server_port": 8080,
  		"db_host": "localhost",
  		"DB_PASSWORD": "super_secret_password",
  		"feature_flags": {
    		"beta_mode": true,
    		"deprecated_v1": false
  	},
  		"optional_webhook": ""
	}`,
)

type FeatureFlags struct {
	BetaMode     bool `json:"beta_mode"`
	DeprecatedV1 bool `json:"deprecated_v1"`
}

type Config struct {
	Port         int          `json:"server_port"`
	DatabaseHost string       `json:"db_host"`
	DatabasePass string       `json:"DB_PASSWORD ,omitempty"`
	FeatureFlags FeatureFlags `json:"feature_flags"`
	Webhook      string       `json:"optional_webhook,omitempty"`
}

func main() {
	var cfg Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	fmt.Printf("Loaded config %+v\n", cfg)
	fmt.Printf("Password %s\n", cfg.DatabasePass)

	cfg.DatabasePass = ""
	output, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		log.Fatalf("Error generating config: %v", err)
	}

	fmt.Println(string(output))
}
