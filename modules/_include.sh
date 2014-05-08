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

# Dump "key=value" formatted params into an associative array.
declare -a PARAMS
for param in "$@"
do
    eqp=$(echo "$param" | grep '=')
    if [ $? -eq 0 ] && [ ! -z "$eqp" ]
    then
	key=$(echo "$param" | cut -d'=' -f1)
	val=$(echo "$param" | cut -d'=' -f2)
	PARAMS["$key"]="$val"
    fi
done

# The getparam function looks at the first argument, and returns the value of
# the "key=value" formatted parameter that was passed to the module.
getparam()
{
    echo ${PARAMS["$1"]}
}

# netmask2cidr converts an IPv4 netmask to CIDR notation.
netmask2cidr()
{
    cidr=0
    IFS=.
    for dec in $1
    do
	case $dec in
	    255) let cidr+=8 ;;
	    254) let cidr+=7 ;;
	    252) let cidr+=6 ;;
	    248) let cidr+=5 ;;
	    240) let cidr+=4 ;;
	    224) let cidr+=3 ;;
	    192) let cidr+=2 ;;
	    128) let cidr+=1 ;;
	    0) ;;
	    *) echo "$dec is not recognized"; exit 1
	esac
    done
    echo "$cidr"
}

# This is the variable that will hold all of the detected system information.
declare -a GOVERN
