package main

// The docopt-friendly USAGE message for argument parsing.
const USAGE = `Govern your infrastructure.

Usage:
  govern facts [--output=(yaml|json)]
  govern [--path=<path>] [--verbose] check
  govern [--path=<path>] [--verbose] play <playbook> [-i <inventory> | --inventory=<inventory>]
  govern -h | --help
  govern --version

Options:
  -h --help                Show this screen
  --inventory=<inventory>  The inventory path, or URI
  --output=(yaml|json)     The output format for non-logging messages [default: json]
  --path=<path>            Specify the directory to look in for modules, tasks, etc.
  --version                Show version and exit
  -v --verbose             Verbose output
`
