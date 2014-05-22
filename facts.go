package main

import (
	"errors"
	"os/exec"
	"strings"
)

type (
	Facts struct {
		Hostname   string
		FQDN       string
		DomainName string
		OS         OperatingSystemFacts
		Net        NetworkFacts
	}

	OperatingSystemFacts struct {
		Version string
		Arch    string
		Name    string
	}

	NetworkFacts struct {
		Interfaces []NetworkInterfaceFacts
	}

	NetworkInterfaceFacts struct {
		Name string
		Addr []string
		MAC  string
	}
)

func GatherFacts() (facts *Facts, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(e.(string))
		}
	}()

	osFacts := OperatingSystemFacts{
		Version: GetFact(CmdOSVersion),
		Arch:    GetFact(CmdOSArch),
		Name:    GetFact(CmdOSName)}

	facts = &Facts{
		Hostname:   GetFact(CmdHostname),
		FQDN:       GetFact(CmdFQDN),
		DomainName: "",
		OS:         osFacts}
	return
}

func GetFact(cmd *exec.Cmd) (output string) {
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	output = strings.Trim(string(out), "\n")
	return
}
