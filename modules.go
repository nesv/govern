package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	ModulePaths []string
)

func init() {
	ModulePaths = make([]string, 0)
	ModulePaths = append(ModulePaths, "/usr/local/etc/govern/modules")
	ModulePaths = append(ModulePaths, "/etc/govern/modules")

	userModsPath := filepath.Join(os.Getenv("HOME"), ".govern", "modules")
	ModulePaths = append(ModulePaths, userModsPath)
}

type Module struct {
	Name           string
	OSFamily       string
	OSVersionMajor string
	OSVersionMinor string
	Path           string
}

func GatherModules() ([]Module, error) {
	mods := make([]Module, 0)
	for _, pth := range ModulePaths {
		if fi, err := os.Stat(pth); err != nil && os.IsNotExist(err) {
			// The error is just because the path does not exist.
			// Carry on.
			continue
		} else if err != nil {
			return nil, err
		} else if !fi.IsDir() {
			return nil, fmt.Errorf("%s is not a directory")
		}

		m, err := gatherModulesInDir(pth)
		if err != nil {
			return nil, err
		}

		mods = append(mods, m...)
	}

	return mods, nil
}

func gatherModulesInDir(pth string) ([]Module, error) {
	matches, err := filepath.Glob(filepath.Join(pth, "*.sh"))
	if err != nil {
		return nil, err
	}

	mods := make([]Module, 0)
	for _, match := range matches {
		parts := strings.SplitN(filepath.Base(match), "_", 4)

		var name, osf, osvmj, osvmn string

		switch nparts := len(parts); nparts {
		case 0:
			return nil, fmt.Errorf("poorly named module %q", match)

		case 1:
			name = strings.TrimRight(parts[0], ".sh")

		case 2:
			name = parts[0]
			osf = strings.TrimRight(parts[1], ".sh")

		case 3:
			name = parts[0]
			osf = parts[1]
			osvmj = strings.TrimRight(parts[2], ".sh")

		case 4:
			name = parts[0]
			osf = parts[1]
			osvmj = parts[2]
			osvmn = strings.TrimRight(parts[3], ".sh")
		}

		m := Module{
			Name:           name,
			OSFamily:       osf,
			OSVersionMajor: osvmj,
			OSVersionMinor: osvmn,
			Path:           match}

		mods = append(mods, m)
	}

	return mods, nil
}
