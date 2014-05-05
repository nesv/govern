#!/usr/bin/env bash
#
# _Darwin.sh
#
# Gathers information about Darwin (OS X) hosts.
#
GOVERN_ARCH=$(uname -m)
GOVERN_OS_VERSION=$(uname -v)
GOVERN_FQDN=$(uname -n)
GOVERN_HOSTNAME=$(hostname -s)
GOVERN_DOMAIN=""
GOVERN_INTERFACES=$(ifconfig | egrep '^[a-z0-9]+:' | awk -F: '{print $1; }')
