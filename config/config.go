package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type IdentifyConfig struct {
	UserAgent string `yaml:"user_agent"`
	OS        string `yaml:"os"`
	Browser   string `yaml:"browser"`
}

type Config struct {
	Mouse         bool           `yaml:"mouse"`
	Timestamps    bool           `yaml:"timestamps"`
	MessagesLimit uint           `yaml:"messages_limit"`
	Identify      IdentifyConfig `yaml:"identify"`
}

func New() *Config {
	return &Config{
		Mouse:         true,
		Timestamps:    false,
		MessagesLimit: 50,
		Identify: IdentifyConfig{
			UserAgent: userAgent,
			OS:        oss,
			Browser:   browser,
		},
	}
}

func (c *Config) Load(path string) error {
	// Create directories mentioned in the path (excluding the name of the file at the end) that do not already exist recursively.
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	// If the configuration file does not exist already, create a new file; otherwise, open the existing file with read-write flag.
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	s, err := f.Stat()
	if err != nil {
		return err
	}

	// If the file is empty, that is, the size of the file is zero, write the default configuration to the file.
	if s.Size() == 0 {
		e := yaml.NewEncoder(f)
		return e.Encode(c)
	}

	return yaml.NewDecoder(f).Decode(&c)
}

func DefaultPath() string {
	path, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	path = filepath.Join(path, Name, "config.yaml")
	return path
}
