package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"gopkg.in/yaml.v1"
)

// Function PrettyPrint prints the provided value v in the specified format.
func PrettyPrint(v interface{}, format string) {
	var p []byte
	var err error
	switch format {
	case "yaml":
		p, err = yaml.Marshal(v)
	case "json":
		p, err = json.MarshalIndent(v, "", "\t")
	default:
		glog.Warningf("unsupported output format: %q", format)
		return
	}
	if err != nil {
		glog.Errorln(err)
		return
	}
	fmt.Printf("%s", p)
}
