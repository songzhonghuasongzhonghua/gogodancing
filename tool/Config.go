package tool

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var _cfg = new(Config)

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(_cfg); err != nil {
		return nil, err
	}

	return _cfg, nil

}
