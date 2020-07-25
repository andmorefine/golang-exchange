package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// ConfigList struct
type ConfigList struct {
	LogFile string

	DbName    string
	SQLDriver string

	Port int
}

// Config 設定
var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		LogFile:   cfg.Section("io").Key("log_file").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Port:      cfg.Section("web").Key("port").MustInt(),
	}
}
