package main

import (
	"errors"
	"net"
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
		Interfaces []NetworkInterfaceFact
	}

	NetworkInterfaceFact struct {
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

	// Gather operating system facts.
	osFacts := OperatingSystemFacts{
		Version: GetFact(CmdOSVersion),
		Arch:    GetFact(CmdOSArch),
		Name:    GetFact(CmdOSName)}

	// Gather networking facts.
	var interfaces []net.Interface
	interfaces, err = net.Interfaces()
	if err != nil {
		return
	}
	ifaceFacts := make([]NetworkInterfaceFact, len(interfaces))
	for i, iface := range interfaces {
		fact := NetworkInterfaceFact{
			Name: iface.Name,
			MAC:  iface.HardwareAddr.String()}

		// Get all of the addresses for the network interface.
		var addrs []net.Addr
		if addrs, err = iface.Addrs(); err != nil {
			return
		}
		ifaddrs := make([]string, len(addrs))
		for _, addr := range addrs {
			ifaddrs = append(ifaddrs, addr.String())
		}
		fact.Addr = ifaddrs

		ifaceFacts[i] = fact
	}

	// Put it all together, and whaddaya get?
	facts = &Facts{
		Hostname:   GetFact(CmdHostname),
		FQDN:       GetFact(CmdFQDN),
		DomainName: "",
		OS:         osFacts,
		Net:        NetworkFacts{Interfaces: ifaceFacts}}
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
