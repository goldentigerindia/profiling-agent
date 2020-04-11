package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)
var ApplicationConfig *Config=new(Config)
type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	FileSystem struct {
		HostProc string `yaml:"hostProc", envconfig:"HOST_PROC"`
		HostSys string `yaml:"hostSys", envconfig:"HOST_SYS"`
		HostEtc string `yaml:"hostEtc", envconfig:"HOST_ETC"`
	} `yaml:"fileSystem"`
}
func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
