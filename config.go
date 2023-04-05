package main

import (
	"encoding/json"
	"flag"
	"os"
)

func (args *Args) ReadFlags() {
	flag.StringVar(&args.Path, "config", "./config.json", "Path to config file")
	flag.Parse()
}

func (config *Config) ReadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	return nil
}
