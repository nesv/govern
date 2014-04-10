package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

type InventoryGroup struct {
	Name     string   `yaml:"group"`
	Hosts    []string `yaml:"hosts,omitempty"`
	Children []string `yaml:"children,omitempty"`
}

func (g InventoryGroup) Check() error {
	// Make sure there are either hosts, or child groups specified.
	if len(g.Hosts) == 0 && len(g.Children) == 0 {
		return fmt.Errorf("no hosts or child groups specified in group")
	}
	return nil
}

func LoadInventoryFile(pth string) ([]InventoryGroup, error) {
	b, err := ioutil.ReadFile(pth)
	if err != nil {
		return nil, err
	}

	var inv []InventoryGroup
	err = yaml.Unmarshal(b, &inv)
	if err != nil {
		return nil, err
	}

	return inv, err
}
