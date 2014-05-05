package main

import (
	"flag"
	"log"
	"os"

	"github.com/golang/glog"
)

func main() {
	PlaybookFile := flag.String("play", "site.yml", "Path to the playbook to execute")
	InventoryFile := flag.String("i", "hosts", "Path to the inventory file")
	CheckAndQuit := flag.Bool("check", false, "Check and exit without running the play")
	DataDir := flag.String("path", "", "Alternate path for handlers, tasks, etc.")
	flag.Parse()

	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	inv, err := LoadInventoryFile(*InventoryFile)
	if err != nil {
		glog.Fatalf("error loading inventory file %q reason=%s", *InventoryFile, err.Error())
	}
	for _, group := range inv {
		glog.V(1).Infof("loaded host group %q", group.Name)
	}

	// Run a sanity check on the inventory groups.
	for _, g := range inv {
		if err = g.Check(); err != nil {
			glog.Fatalf("error in group %q: %s", g.Name, err.Error())
		}
	}

	// Gather a list of the available modules.
	modules, err := GatherModules()
	if err != nil {
		glog.Fatalf("Error gathering modules: %s", err.Error())
	}
	if len(modules) == 0 {
		glog.Fatalln("no modules loaded")
	}
	for _, module := range modules {
		glog.V(1).Infof("loaded module %q", module.Name)
	}

	// Build a list of paths for things like handlers, tasks, etc., from
	// the DataDir flag ("-path") value.
	var basedir string
	if *DataDir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			glog.Fatalln(err)
		}
		basedir = pwd
	} else {
		basedir = *DataDir
	}

	// Load handlers.
	handlers, err := LoadHandlers(basedir)
	if err != nil {
		glog.Fatalln(err.Error())
	}
	for name, _ := range handlers {
		glog.V(1).Infof("loaded handler %q", name)
	}

	// Load tasks.
	tasks, err := LoadTasks(basedir)
	if err != nil {
		glog.Fatalln(err)
	}
	for _, task := range tasks {
		if err := task.Check(&modules); err != nil {
			glog.Fatalf("error with task %q from file %q: %v", task.Name, task.file, err)
		}
		glog.V(1).Infof("loaded task %q", task.Name)
	}

	// Load the roles.

	// Load the playbook.
	plays, err := LoadPlaybook(*PlaybookFile)
	if err != nil {
		glog.Fatalf("error loading playbook %q: %s", *PlaybookFile, err.Error())
	}
	for _, p := range plays {
		// Check the plays.
		if err := p.Check(); err != nil {
			glog.Fatalf("error in play %q: %s", p.Name, err.Error())
		}
		glog.V(1).Infof("loaded play %q", p.Name)
	}

	// Bail out here if the user wanted only to check the format of their
	// plays, roles, tasks, etc.
	if *CheckAndQuit {
		glog.Infoln("Checks passed")
		os.Exit(0)
	}

	return
}
