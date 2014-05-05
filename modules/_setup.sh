#!/usr/bin/env bash
#
# _setup.sh
#
# All-encompassing Bash file to include, which includes other Bash scripts that
# detect operating system information, network interface information, and more.
#
# NOTE:
#    "_osdetect.sh" must be first!
#

. _osdetect.sh
[ -z "$GOVERN_OS" ] && echo "could not detect operating system" && exit 1

# Check to make sure the OS-specific detection script exists.
scr="_${GOVERN_OS.sh}"
[ ! -e "$scr" ] && echo "unsupported operating system: ${GOVERN_OS}" && exit 1

. "_${GOVERN_OS}.sh"
