SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

sudo rm -rf $SCRIPT_DIR/{config,data,logs}
