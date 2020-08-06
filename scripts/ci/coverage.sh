#!/bin/bash

echo "Checking variables"
echo $GITHUB_ACTIONS
echo $GITHUB_REF

# Checks if the tag script is being run on Github Actions
if [ ! $GITHUB_ACTIONS ]; then
    echo "This script needs to be run on Github Actions platform"
    exit 1
fi

if [ ! "$GITHUB_REF" == "master" ]; then
    echo "This build is not in the master branch"
    exit 1
fi

# Checks is the CODECOV_TOKEN token has been set
if [ -z $CODECOV_TOKEN ]; then
    echo "CODECOV_TOKEN has not been set!"
    exit 2
fi

# Checks if the coverage file exists from the makefile
if [ ! -f coverage.txt ]; then
    echo "coverage.txt file not found!"
    exit 2
fi

bash <(curl -s https://codecov.io/bash)

exit 1