package main

import (
	"flag"

	"github.com/golang/glog"
)

var (
	PlaybookFile  = flag.String("play", "site.yml", "Path to the playbook to execute")
	InventoryFile = flag.String("i", "hosts", "Path to the inventory file")
	LimitHosts    = flag.String("l", "", "Limit hosts")
)

func main() {
	flag.Parse()

	inv, err := LoadInventoryFile(*InventoryFile)
	if err != nil {
		glog.Fatalf("error loading inventory file %q reason=%s", *InventoryFile, err.Error())
	}

	if ngroups := len(inv); ngroups == 1 {
		glog.V(1).Info("Loaded 1 group from inventory")
	} else {
		glog.V(1).Infof("Loaded %d groups from inventory", ngroups)
	}

	// Run a sanity check on the inventory groups.
	for _, g := range inv {
		if err = g.Check(); err != nil {
			glog.Fatalf("Error in group %q: %s", g.Name, err.Error())
		}
	}

	return
}
