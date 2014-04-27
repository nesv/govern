package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

func LoadYAML(path string, v interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(b, v)
	return err
}
