package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Handler struct {
	Name       string
	Tags       []string
	module     string
	moduleArgs string
}

func (h Handler) Check(m *[]Module) error {
	if h.Name == "" {
		return fmt.Errorf("handler name is empty")
	}

	for _, mod := range *m {
		if h.module == mod.Name {
			return nil
		}
	}
	return fmt.Errorf("no such module %q", h.module)
}

func (h Handler) ModuleName() string {
	return h.module
}

func (h Handler) ModuleArgs() string {
	return h.moduleArgs
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
		var y RawYAML
		y, err = LoadYAML(m)
		if err != nil {
			return
		}

		// Pull values out of the map, and create a handler struct
		// from the values.
		var h Handler
		h, err = NewHandlerFromRawYAML(&y)
		if err != nil {
			return
		}

		handlers = append(handlers, h)
	}

	return
}

// Function NewHandlerFromMap creates a new Handler from map m.
//
// This function pulls out reserved keys from the map, deletes them from the
// map, then iterates over the remaining keys trying to match them to names
// of loaded modules.
func NewHandlerFromRawYAML(y *RawYAML) (Handler, error) {
	h := Handler{}
	if name := y.ToString("name"); name == "" {
		return h, fmt.Errorf("handler name must not be empty")
	} else {
		h.Name = name
	}

	// Check for any, optional tags.
	if tags := y.ToStringSlice("tags"); tags != nil {
		h.Tags = tags
	}

	// At this point, we will just take the first, left-over key from the
	// map, and assume that is the module we want. However, we will also
	// want to check that we have any keys left in the map.
	if len(*y) == 0 {
		return h, fmt.Errorf("no module specified in handler %q", h.Name)
	}
	for k, v := range *y {
		h.module = k
		if s, ok := v.(string); ok {
			h.moduleArgs = s
		}
	}

	return h, nil
}
