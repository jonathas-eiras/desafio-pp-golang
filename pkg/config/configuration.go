package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

func GetApplication() Application {

	var application Application

	filename, _ := filepath.Abs("./../../resources/application.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &application)
	if err != nil {
		panic(err)
	}

	return application
}
