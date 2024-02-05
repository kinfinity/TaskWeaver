package cli_runtime

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ClientOptions struct {
	Name       string
	Version    string
	ConfigFile string
	logger     *log.Logger
	LogFile    string
}

// get Logger
func (co *ClientOptions) GetLogger() *log.Logger {
	if co.logger == nil {
		return log.Default()
	}
	return co.logger
}

// set Logger
func (co *ClientOptions) SetLogger(logger *log.Logger) {
	co.logger = logger
}

type WeaverConfig struct {
	ClientName  string `yaml:"client-name"`
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log-level"`
	LogFile     string `yaml:"log-file"`
}

type Config struct {
	WeaverConfig WeaverConfig `yaml:"weaver_config"`
}

func InitializeConfig(logger *log.Logger, configDir string, configFilePath string) (err error) {
	// Create the directory if it doesn't exist
	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		logger.Fatalln("Error creating directory:", err)
	}
	homeDir, _ := os.UserHomeDir()
	// create default config
	weaverConfig := WeaverConfig{
		ClientName:  "weaver",
		Environment: "",
		LogLevel:    "info",
		LogFile:     homeDir + "/.weaver/config/weaver.log",
	}
	config := Config{
		WeaverConfig: weaverConfig,
	}

	// Marshal the configuration to YAML
	configYAML, err := yaml.Marshal(&config)
	if err != nil {
		logger.Fatalln("Error marshaling configuration to YAML:", err)
	}
	// Write the YAML to the config file
	err = os.WriteFile(configFilePath, configYAML, 0600)
	if err != nil {
		logger.Fatalln("Error writing configuration to file:", err)
	}

	// // create default log file
	err = os.WriteFile(weaverConfig.LogFile, []byte{}, 0666)
	if err != nil {
		logger.Fatalln("Error creating default log file:", err)
	}

	return nil
}

// Need specific options for Windows | MAC | Liunx
