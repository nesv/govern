#!/usr/bin/env bash
#
# setup.sh
#
# The setup module gathers information about a host, and dumps the contents.
#
USAGE="setup"
NO_ARGS_REQUIRED="true"

. _include.sh
. _setup.sh

echo "${GOVERN_IPV4[@]}"
