package main

import "os/exec"

var (
	CmdOSVersion = exec.Command("uname", "-v")
	CmdOSArch    = exec.Command("uname", "-m")
	CmdOSName    = exec.Command("uname")
	CmdFQDN      = exec.Command("uname", "-n")
	CmdHostname  = exec.Command("hostname", "-s")
)
