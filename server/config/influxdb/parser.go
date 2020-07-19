package influxdb

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	URL             string        `yaml:"url"`
	OrganizationId  string        `yaml:"organizationId"`
	BucketId        string        `yaml:"bucketId"`
	Token           string        `yaml:"token"`
	Measurement     string        `yaml:"measurement"`
	Interval        time.Duration `yaml:"interval"`
	AlignTimestamps bool          `yaml:"alignTimestamps"`
}

//ParseConfig parses configuration file and returns corresponding object
func ParseConfig(configPath string) *Config {
	path, err := filepath.Abs(configPath)
	if err != nil {
		log.Fatalf("did not get the absolute path: %v", err)
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("could not parse: %v", err)
	}

	return &config
}
