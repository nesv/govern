## _include.sh
##
## A generic include that provides some helper functions for modules.
##

# Check to make sure the module has a USAGE variable defined.
if [ -z "$USAGE" ]
then
    echo "USAGE is not defined"
    exit 255
fi

# Check to see if any arguments were provided, and fail out if the module
# has not set the "NO_ARGS_REQUIRED=true|1|whatever" variable.

if [ "$#" == "0" ] && [ "x$NO_ARGS_REQUIRED" == "x" ]
then
    echo "No arguments provided"
    exit 254
fi

PARAMS="$@"

# Dump "key=value" formatted params into an associative array.
declare -a KWPARAMS
for param in "${PARAMS[@]}"
do
    key=$(echo "$param" | cut -d'=' -f1)
    val=$(echo "$param" | cut -d'=' -f2)
    KWPARAMS["$key"]="$val"
done

# The getparam function looks at the first argument, and returns the value of
# the "key=value" formatted parameter that was passed to the module.
getparam()
{
    echo ${KWPARAMS["$1"]}
}
