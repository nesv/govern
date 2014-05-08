#!/usr/bin/env bash
#
# _Linux.sh
#
# Gathers information about Linux hosts.
#
GOVERN_ARCH=$(uname -m)
GOVERN_OS_VERSION=$(uname -v)
GOVERN_FQDN=$(uname -n)
GOVERN_HOSTNAME=$(hostname -s)
GOVERN_DOMAIN=$(hostname -d)

declare -a GOVERN_INTERFACES
declare -a GOVERN_IPV4
declare -a GOVERN_IPV4_CIDR
if [ ! -z "$(which ip 2>/dev/null)" ]
then
    # Use the ip(8) command to get network interface information.
    for iface in $(ip addr | egrep '^\s*inet' | sed 's/^\s*//')
    do
	ifname=$(echo $iface | awk '{ print $NF; }')
	GOVERN_INTERFACES+=("$ifname")

	addr=$(echo $iface | awk '{ print $2; }' | sed 's/\/[0-9]*$//')
	GOVERN_IPV4[$ifname]=$addr

	cidr=$(echo $iface | awk '{ print $2; }' | sed 's/^[.0-9]*\///')
	GOVERN_IPV4_CIDR[$ifname]=$cidr
    done
else
    # Fall back to using ifconfig(8).
    for ifname in $(ifconfig -s | sed '1d' | awk '{ print $1; }')
    do
	GOVERN_INTERFACES+=("$ifname")
	
	ifinfo=$(ifconfig $ifname | egrep '^\s*inet\s' | sed 's/^\s*//')
	addr=$(echo $ifinfo | awk '{ print $2; }' | cut -d: -f2)
	GOVERN_IPV4[$ifname]=$addr

	netmask=$(echo $ifinfo | awk '{ print $4; }' | cut -d: -f2) 
	GOVERN_IPV4_CIDR[$ifname]=$(netmask2cidr $netmask)
    done
fi
