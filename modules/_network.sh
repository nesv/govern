#!/usr/bin/env bash
#
# _network.sh
#
# Detects networking information.
#

set -e
set -x

## Get the FQDN, short hostname, and domain name for the system.
declare _govern_net_FQDNCmd
declare _govern_net_HostnameCmd
declare _govern_net_DomainCmd
case "$GOVERN_OS" in
    OpenBSD | Darwin )
	_govern_net_FQDNCmd="uname -n"
	_govern_net_HostnameCmd="hostname -s"
	_govern_net_DomainCmd="true"
	;;

    Linux )
	_govern_net_FQDNCmd="uname -n"
	_govern_net_HostnameCmd="hostname -s"
	_govern_net_DomainCmd="hostname -d"
	;;
esac
GOVERN_FQDN=$(${_govern_net_FQDNCmd})
GOVERN_HOSTNAME=$(${_govern_net_HostnameCmd})
GOVERN_DOMAIN=$(${_govern_net_DomainCmd})

## Get a list of the network interfaces.
function get_network_interfaces_Darwin {
    ifconfig 2>&1 | egrep '^[a-z0-9]+:' | awk -F: '{ print $1; }'
}

function get_network_interfaces_OpenBSD {
    ifconfig 2>&1 | egrep '^[a-z0-9]+:' | awk -F: '{ print $1; }'
}

function get_network_interfaces_Linux {
    ip addr | egrep '^[1-9]+:' | awk -F: '{print $2; }' | sed 's/\s//'
}

GOVERN_INTERFACES=$(eval "get_network_interfaces_${GOVERN_OS}")

# declare GOVERN_INTERFACES
# case "$GOVERN_OS" in
#     Darwin | OpenBSD )
# 	_govern_net_ListInterfacesCmd="ifconfig 2>&1 | egrep '^[a-z0-9]+:' | awk -F: '{ print $1; }'"
# 	;;

#     Linux )
# 	_govern_net_ListInterfacesCmd="ip addr | egrep '^[1-9]+:' | awk -F: '{ print $2; }' | sed -e 's/\s//'"
# 	;;

#     * )
# 	echo "unsupported operating system"
# 	exit 1
# 	;;
# esac
# GOVERN_INTERFACES=$(eval "$_govern_net_ListInterfacesCmd")
