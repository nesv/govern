#!/usr/bin/env bash

USAGE="echo MESSAGE"

. _include.sh

message=$(getparam "message")
[ ! -z "$message" ] && echo "$message" || echo "${PARAMS[@]}"
