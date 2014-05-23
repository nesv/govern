package main

import (
	"fmt"
	"log"
	"os"

	"github.com/docopt/docopt-go"
)

const VERSION = "0.0.1"

var (
	Verbose       bool   = false
	InventoryFile string = "hosts"
	Playbook      string = "site.yml"
	OutputFormat  string
)

func main() {
	log.SetFlags(0)
	args, err := docopt.Parse(USAGE, nil, true, VERSION, false)
	if err != nil {
		log.Fatalln(err)
	}

	if verbose, ok := args["--verbose"].(bool); ok {
		Verbose = verbose
	}

	if showVersion, ok := args["--version"].(bool); showVersion && ok {
		fmt.Println(VERSION)
		return
	}

	// Gather facts about the current system.
	facts, err := GatherFacts()
	if err != nil {
		log.Fatalf("failed to gather facts: %v", err.Error())
	}
	if outputfmt, ok := args["--output"].(string); ok {
		OutputFormat = outputfmt
	}
	if showFacts, ok := args["facts"].(bool); showFacts && ok {
		PrettyPrint(facts, OutputFormat)
		return
	}

	// Load the inventory file.
	InventoryFile, ok := args["--inventory"].(string)
	if !ok {
		log.Fatalln("could not coerce inventory to a string")
	}
	inv, err := LoadInventoryFile(InventoryFile)
	if err != nil {
		log.Fatalf("error loading inventory file %q reason=%s", InventoryFile, err.Error())
	}
	for _, group := range inv {
		log.Printf("loaded host group %q", group.Name)
	}

	// Run a sanity check on the inventory groups.
	for _, g := range inv {
		if err = g.Check(); err != nil {
			log.Fatalf("error in group %q: %s", g.Name, err.Error())
		}
	}

	// Gather a list of the available modules.
	modules, err := GatherModules()
	if err != nil {
		log.Fatalf("Error gathering modules: %s", err.Error())
	}
	if len(modules) == 0 {
		log.Fatalln("no modules loaded")
	}
	for _, module := range modules {
		if Verbose {
			log.Printf("loaded module %q", module.Name)
		}
	}

	// Build a list of paths for things like handlers, tasks, etc., from
	// the DataDir flag ("--path") value.
	datadir, ok := args["--path"].(string)
	if !ok {
		log.Fatalln("could not coerce --path value to string")
	}
	var basedir string
	if datadir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		basedir = pwd
	} else {
		basedir = datadir
	}

	// Load handlers.
	handlers, err := LoadHandlers(basedir)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for name, _ := range handlers {
		if Verbose {
			log.Printf("loaded handler %q", name)
		}
	}

	// Load tasks.
	tasks, err := LoadTasks(basedir)
	if err != nil {
		log.Fatalln(err)
	}
	for _, task := range tasks {
		if err := task.Check(&modules); err != nil {
			log.Fatalf("error with task %q from file %q: %v", task.Name, task.file, err)
		}
		log.Printf("loaded task %q", task.Name)
	}

	// Load the roles.

	// Load the playbook.
	playbookFile, ok := args["<playbook>"].(string)
	if !ok {
		log.Fatalln("could not coerce <playbook> to a string")
	}
	plays, err := LoadPlaybook(playbookFile)
	if err != nil {
		log.Fatalf("error loading playbook %q: %s", playbookFile, err.Error())
	}
	for _, p := range plays {
		// Check the plays.
		if err := p.Check(); err != nil {
			log.Fatalf("error in play %q: %s", p.Name, err.Error())
		}
		if Verbose {
			log.Printf("loaded play %q", p.Name)
		}
	}

	// Bail out here if the user wanted only to check the format of their
	// plays, roles, tasks, etc.
	if check, ok := args["check"].(bool); ok && check {
		log.Println("checks passed")
		return
	} else if !ok {
		log.Fatalln("could not coerce check to a bool")
	}

	return
}
