package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v1"
)

type Handler struct {
	Name   string `yaml:"name"`
	Action string `yaml:"action"`
}

func (h Handler) Check(m *[]Module) error {
	if h.Name == "" {
		return fmt.Errorf("handler name is empty")
	}
	if h.Action == "" {
		return fmt.Errorf("no action specified for handler %q", h.Name)
	}

	a := strings.Split(h.Action, " ")
	for _, mod := range *m {
		if a[0] == mod.Name {
			return nil
		}
	}
	return fmt.Errorf("no such module %q", a[0])
}

func LoadHandlers(basedir string) (handlers []Handler, err error) {
	// The handlers directory path.
	var hd string
	if basedir == "" {
		var pwd string
		if pwd, err = os.Getwd(); err != nil {
			return
		}
		hd = filepath.Join(pwd, "handlers")
	} else {
		hd = filepath.Join(basedir, "handlers")
	}

	var fi os.FileInfo
	if fi, err = os.Stat(hd); err != nil && os.IsNotExist(err) {
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
	handlers = make([]Handler, 0)

	glob := filepath.Join(hd, "*.yml")
	matches, err := filepath.Glob(glob)
	if err != nil {
		return
	}

	for _, m := range matches {
		b, e := ioutil.ReadFile(m)
		if e != nil {
			err = e
			return
		}

		var h []Handler
		if e = yaml.Unmarshal(b, &h); e != nil {
			err = e
			return
		}

		handlers = append(handlers, h...)
	}

	return
}
