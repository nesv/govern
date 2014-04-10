package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

type Play struct {
	Name   string   `yaml:"play"`
	Hosts  []string `yaml:"hosts"`
	Roles  []string `yaml:"roles,flow"`
	Serial int      `yaml:"serial,omitempty"`
}

func (p Play) Check() error {
	if p.Name == "" {
		return fmt.Errorf("play name cannot be empty")
	}

	if len(p.Hosts) == 0 {
		return fmt.Errorf("no host groups specified")
	}

	if len(p.Roles) == 0 {
		return fmt.Errorf("no roles specified")
	}

	return nil
}

func LoadPlaybook(pth string) ([]Play, error) {
	b, err := ioutil.ReadFile(pth)
	if err != nil {
		return nil, err
	}

	var plays []Play
	err = yaml.Unmarshal(b, &plays)
	return plays, err
}
