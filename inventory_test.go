package main

import "testing"

func TestLoadInventoryFile(t *testing.T) {
	_, err := LoadInventoryFile("testfiles/hosts")
	if err != nil {
		t.Error(err)
	}
}
