package io

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func ExistsFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFolder(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func WriteYamlFile(filename string, payload interface{}) error {
	err := CreateFolder(filepath.Dir(filename))
	if err != nil {
		return err
	}
	yamlData, errorMarshalling := yaml.Marshal(payload)
	if errorMarshalling != nil {
		return errorMarshalling
	}
	return os.WriteFile(filename, yamlData, 0755)
}
