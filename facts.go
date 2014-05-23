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
		Addr NetworkInterfaceAddrsFact
		MAC  string
	}

	NetworkInterfaceAddrsFact struct {
		IPv4 []IPAddrFact
		IPv6 []IPAddrFact
	}

	IPAddrFact struct {
		Addr string
		Mask string
	}
)

// Function GatherFacts is the main function to call for gathering facts on the
// system.
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
	netFacts, err := GatherNetworkingFacts()
	if err != nil {
		return
	}

	// Put it all together, and whaddaya get?
	facts = &Facts{
		Hostname:   GetFact(CmdHostname),
		FQDN:       GetFact(CmdFQDN),
		DomainName: "",
		OS:         osFacts,
		Net:        netFacts}
	return
}

// Function GetFact runs the specified cmd, and returns the output.
//
// This function will panic on any errors.
func GetFact(cmd *exec.Cmd) (output string) {
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	output = strings.Trim(string(out), "\n")
	return
}

// Function GatherNetworkingFacts gathers fact information about networking
// interfaces.
func GatherNetworkingFacts() (facts NetworkFacts, err error) {
	interfaces, err := net.Interfaces()
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
		ipv4 := make([]IPAddrFact, 0)
		ipv6 := make([]IPAddrFact, 0)
		for _, addr := range addrs {
			parts := strings.Split(addr.String(), "/")
			ipaf := IPAddrFact{Addr: parts[0], Mask: parts[1]}
			ip := net.ParseIP(parts[0])
			if ip4 := ip.To4(); ip4 == nil {
				ipv6 = append(ipv6, ipaf)
			} else {
				ipv4 = append(ipv4, ipaf)
			}
		}
		fact.Addr = NetworkInterfaceAddrsFact{IPv4: ipv4, IPv6: ipv6}
		ifaceFacts[i] = fact
	}

	facts = NetworkFacts{Interfaces: ifaceFacts}
	return
}
