package main

import (
	"flag"
	"log"
	"os"
)

var (
	PlaybookFile  = flag.String("play", "site.yml", "Path to the playbook to execute")
	InventoryFile = flag.String("i", "hosts", "Path to the inventory file")
	LimitHosts    = flag.String("l", "", "Limit hosts")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	inv, err := LoadInventoryFile(*InventoryFile)
	if err != nil {
		log.Fatalf("error loading inventory file %q reason=%s", *InventoryFile, err.Error())
	}

	if ngroups := len(inv); ngroups == 1 {
		log.Println("Loaded 1 group from inventory")
	} else {
		log.Printf("Loaded %d groups from inventory", ngroups)
	}

	// Run a sanity check on the inventory groups.
	for _, g := range inv {
		if err = g.Check(); err != nil {
			log.Fatalf("Error in group %q: %s", g.Name, err.Error())
		}
	}

	// Load the playbook.
	plays, err := LoadPlaybook(*PlaybookFile)
	if err != nil {
		log.Fatalf("Error loading playbook %q: %s", *PlaybookFile, err.Error())
	}

	if nplays := len(plays); nplays == 1 {
		log.Println("Loaded 1 play")
	} else {
		log.Printf("Loaded %d plays", len(plays))
	}

	// Check the plays.
	for _, p := range plays {
		if err := p.Check(); err != nil {
			log.Fatalf("Error in play %q: %s", p.Name, err.Error())
		}
	}

	return
}
