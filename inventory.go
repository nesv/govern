package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/glog"
	yaml "gopkg.in/yaml.v1"
)

type Inventory struct {
	Groups []InventoryGroup `yaml:"groups,flow"`
}

func (inv *Inventory) Check() error {
	var err error
	for _, g := range inv.Groups {
		if err = g.Check(); err != nil {
			return err
		}
	}
	return nil
}

type InventoryGroup struct {
	Name     string   `yaml:"group"`
	Hosts    []string `yaml:"hosts,omitempty"`
	Children []string `yaml:"children,omitempty"`
}

func (g *InventoryGroup) Check() error {
	if g.Name == "" {
		return fmt.Errorf("cannot have an empty group name")
	}

	// Make sure there are either hosts, or child groups specified.
	if len(g.Hosts) == 0 && len(g.Children) == 0 {
		return fmt.Errorf("no hosts or child groups specified in group %q", g.Name)
	}
	return nil
}

func LoadInventoryFile(pth string) (*Inventory, error) {
	b, err := ioutil.ReadFile(pth)
	if err != nil {
		return nil, err
	}

	glog.V(4).Infof("Read %d bytes from %s", len(b), pth)

	var inv Inventory
	err = yaml.Unmarshal(b, &inv)
	if err != nil {
		return nil, err
	}

	return &inv, err
}
