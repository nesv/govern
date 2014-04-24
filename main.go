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
	CheckAndQuit  = flag.Bool("check", false, "Check and exit without running the play")
	DataDir       = flag.String("d", "", "Alternate path for handlers, tasks, etc.")
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
		log.Println("1 host group loaded from inventory")
	} else {
		log.Printf("%d host groups loaded from inventory", ngroups)
	}

	// Run a sanity check on the inventory groups.
	for _, g := range inv {
		if err = g.Check(); err != nil {
			log.Fatalf("Error in group %q: %s", g.Name, err.Error())
		}
	}

	// Gather a list of the available modules.
	mods, err := GatherModules()
	if err != nil {
		log.Fatalf("Error gathering modules: %s", err.Error())
	}

	if nmods := len(mods); nmods == 0 {
		log.Fatalln("No modules found")
	} else if nmods == 1 {
		log.Println("1 module loaded")
	} else {
		log.Printf("%d modules loaded", nmods)
	}

	// Load handlers.

	// Load tasks.

	// Load the roles.

	// Load the playbook.
	plays, err := LoadPlaybook(*PlaybookFile)
	if err != nil {
		log.Fatalf("Error loading playbook %q: %s", *PlaybookFile, err.Error())
	}

	if nplays := len(plays); nplays == 1 {
		log.Println("1 play loaded")
	} else {
		log.Printf("%d plays loaded", len(plays))
	}

	// Check the plays.
	for _, p := range plays {
		if err := p.Check(); err != nil {
			log.Fatalf("Error in play %q: %s", p.Name, err.Error())
		}
	}

	// Bail out here if the user wanted only to check the format of their
	// plays, roles, tasks, etc.
	if CheckAndQuit {
		log.Println("Checks passed")
		os.Exit(0)
	}

	return
}
