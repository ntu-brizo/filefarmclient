package config

import (
	"encoding/json"
	"os"
)

const configPath string = "/config/conf.json"

// FarmerRecord keeps the connection information of a farmer
type FarmerRecord struct {
	RealIP string `json:"realIp"`
	IoPort int32  `json:"ioPort"`
}

// Config is the structure which maps config json
type Config struct {
	Network struct {
		ConnectionTimeout int `json:"connectionTimeout"`
		StreamTimeout     int `json:"streamTimeout"`
		RetryTimeout      int `json:"retryTimeout"`
		MaxRetryTimes     int `json:"maxRetryTimes"`
	}
	Shard struct {
		NumShards          int `json:"numShards"`
		NumRedundantShards int `json:"numRedundantShards"`
	}
	Farmers []FarmerRecord
}

// Parse parses config json and return a Config struct
func Parse() (config Config, err error) {
	pwd, _ := os.Getwd()
	configFile, err := os.Open(pwd + configPath)
	defer configFile.Close()
	if err != nil {
		return
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}
