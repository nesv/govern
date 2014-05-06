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

declare -a GOVERN_INTERFACES
for iface in $(ifconfig | egrep '^[a-z0-9]+:' | awk -F: '{ print $1; }')
do
    GOVERN_INTERFACES+=("$iface")
done

# For each interface, get the IPv4 addresses.
declare -a GOVERN_IPV4
for iface in "${GOVERN_INTERFACES[@]}"
do
    addr=$(ifconfig $iface | egrep 'inet\s' | awk '{ print $2; }')
    GOVERN_IPV4[$iface]=$addr
done
