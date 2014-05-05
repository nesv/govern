#!/usr/bin/env bash

USAGE="package name=PACKAGE_NAME state=<installed|latest|absent>"

. _include.sh

pkgname=$(getparam "name")
[ -z "$pkgname" ] && echo "no package name specified" && exit 1

pkgstate=$(getparam "state")
[ -z "$pkgstate" ] && echo "no package state specified" && exit 1

