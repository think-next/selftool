package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var _config Config

type Config struct {
	Env struct {
		GoPath      string `yaml:"go_path"`
		ProjectPath string `yaml:"project_path"`
	} `yaml:"env"`
}

func init() {
	file, err := os.Open("./config.yaml")
	if err != nil {
		return
	}
	defer file.Close()

	// load config
	data, _ := ioutil.ReadAll(file)
	if err := yaml.Unmarshal(data, &_config); err != nil {
		panic(err)
	}
}

func GetProjectPath() string {
	return _config.Env.GoPath + "/src/" + _config.Env.ProjectPath
}

func GetGoPath() string {
	return _config.Env.GoPath
}
