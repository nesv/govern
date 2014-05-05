#!/usr/bin/env
#
# _osdetect.sh
#
# Detects various portions of the operating system we are running on.
#

set -e
set -x

## Get the name of the operating system.
GOVERN_OS=$(uname)
