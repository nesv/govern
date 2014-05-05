package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type (
	Task struct {
		Name   string   `yaml:"name"`
		Module string   `yaml:"module"`
		Args   string   `yaml:"args"`
		Notify []string `yaml:"notify,flow"`
		file   string
	}
)

func (t *Task) Check(modules *[]Module) error {
	if len(t.Name) == 0 {
		return fmt.Errorf("task name cannot be empty")
	}
	if len(t.Module) == 0 {
		return fmt.Errorf("task must name a module")
	} else {
		for _, module := range *modules {
			if t.Module == module.Name {
				return nil
			}
		}
		return fmt.Errorf("no such module %q", t.Module)
	}
	return nil
}

func LoadTasks(basedir string) (tasks map[string]*Task, err error) {
	tasksDir := filepath.Join(basedir, "tasks")
	var fi os.FileInfo
	if fi, err = os.Stat(tasksDir); err != nil && os.IsNotExist(err) {
		err = fmt.Errorf("no such directory %q", tasksDir)
		return
	} else if err != nil {
		return
	} else if !fi.IsDir() {
		err = fmt.Errorf("%q is not a directory", tasksDir)
		return
	}

	// Load all of ze YAMLs!
	tasks = make(map[string]*Task, 0)
	glob := filepath.Join(tasksDir, "*.yml")
	matches, err := filepath.Glob(glob)
	if err != nil {
		return
	}

	for _, match := range matches {
		var t []Task
		if err = LoadYAMLFileInto(match, &t); err != nil {
			return
		}
		for _, task := range t {
			task.file = match
			if tt, exists := tasks[task.Name]; exists {
				err = fmt.Errorf("task %q in file %q already exists; previously defined in %q", task.Name, task.file, tt.file)
				return
			}
			tasks[task.Name] = &task
		}
	}
	return
}
