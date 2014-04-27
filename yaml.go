package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

type (
	RawYAML map[string]interface{}
)

func LoadYAML(path string) (RawYAML, error) {
	var y RawYAML
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return y, err
	}

	err = yaml.Unmarshal(b, &y)
	return y, err
}

func (y RawYAML) ToString(key string) string {
	if v, found := y[key]; found {
		if s, ok := v.(string); ok {
			delete(y, key)
			return s
		}
	}
	return ""
}

func (y RawYAML) ToStringSlice(key string) []string {
	if v, found := y[key]; found {
		if s, ok := v.([]string); ok {
			delete(y, key)
			return s
		}
	}
	return nil
}

func (y RawYAML) Length() int {
	return len(y)
}
