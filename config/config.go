package config

import (
	"github.com/spf13/viper"
	"log"
)

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type Config struct {
	Port     string         `json:"port"`
	Database DatabaseConfig `json:"database"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("json")   // config file type
	viper.AddConfigPath(".")      // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &cfg, nil
}

//// LoadConfig loads configuration from the given JSON file path.
//func LoadConfig(path string) (*Config, error) {
//	file, err := os.Open(path)
//	if err != nil {
//		return nil, fmt.Errorf("could not open config file: %w", err)
//	}
//	defer file.Close()
//
//	var config Config
//	decoder := json.NewDecoder(file)
//	if err := decoder.Decode(&config); err != nil {
//		return nil, fmt.Errorf("could not decode config JSON: %w", err)
//	}
//
//	return &config, nil
//}
