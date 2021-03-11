package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration returns struct
type Configuration struct {
	BatchSize              int
	RequestTTLInMinutes	   int
	MaxActiveRequestsCount int
	MaxRequestByteSize     int64
}

// GetConfiguration returns Configuration
func GetConfiguration(environment string) (Configuration, error) {
	file, err := os.Open("api/config/config_" + environment + ".json")
	if err != nil {
		fmt.Println("error:", err)
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	defer file.Close()
	return configuration, err
}
