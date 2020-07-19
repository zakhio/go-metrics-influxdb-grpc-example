package influxdb

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"time"
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

func ParseConfig(filename string) *Config {
	path, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
