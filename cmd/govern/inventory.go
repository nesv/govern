package main

type Inventory []InventoryGroup

type InventoryGroup struct {
	Name  string          `yaml:"name"`
	Hosts []InventoryHost `yaml:"hosts"`
}

type InventoryHost struct {
	Name    string `yaml:"name"`
	SSHHost string `yaml:"ssh_host"`
	SSHUser string `yaml:"ssh_user"`
}

func loadInventory(pth string) (*Inventory, error) {

	return nil, nil
}
