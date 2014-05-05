package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type (
	Handler struct {
		Name   string   `yaml:"name"`
		Tags   []string `yaml:"tags,flow"`
		Module string   `yaml:"module"`
		Args   string   `yaml:"args"`
		file   string
	}
)

func LoadHandlers(basedir string) (handlers map[string]Handler, err error) {
	var handlersDir string = filepath.Join(basedir, "handlers")
	var fi os.FileInfo
	if fi, err = os.Stat(handlersDir); err != nil && os.IsNotExist(err) {
		// There is no handlers subdirectory.
		err = fmt.Errorf("no subdirectory %q", "handlers")
		return
	} else if err != nil {
		return
	}

	// Check to make sure $PWD/handlers is a directory.
	if !fi.IsDir() {
		err = fmt.Errorf("%q is not a directory", "handlers")
		return
	}

	// Load all of the YAML files in the directory, and build up the list
	// of handlers.
	handlers = make(map[string]Handler, 0)

	glob := filepath.Join(handlersDir, "*.yml")
	matches, err := filepath.Glob(glob)
	if err != nil {
		return
	}

	for _, m := range matches {
		var h []Handler
		if err = LoadYAMLFileInto(m, &h); err != nil {
			return
		}
		for _, v := range h {
			v.file = m
			if hh, exists := handlers[v.Name]; exists {
				// There is already a handler with that name
				// that has been loaded. This is a conflict.
				err = fmt.Errorf("handler %q in file %q already exists; previously defined in %q", v.Name, v.file, hh.file)
				return
			}
			handlers[v.Name] = v
		}
	}

	return
}
