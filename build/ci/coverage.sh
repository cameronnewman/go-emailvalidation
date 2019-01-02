#!/bin/bash

echo "Checking variables"
echo $TRAVIS
echo $TRAVIS_TAG
echo $TRAVIS_BRANCH
echo $TRAVIS_PULL_REQUEST_BRANCH
echo $TRAVIS_PULL_REQUEST

# Checks if the tag script is being run on TravisCI
if [ ! $TRAVIS ]; then
    echo "This script needs to be run on TravisCI platform"
    exit 1
fi

# Checks if the build was triggered by a tag commit.
if [ ! -z "$TRAVIS_TAG" ]; then
    echo "This build was triggered by a git tag push"
    exit 1
fi

if [ "$TRAVIS_PULL_REQUEST" -eq "$TRAVIS_PULL_REQUEST" 2>/dev/null ] || [[ "$TRAVIS_PULL_REQUEST" =~ ^[0-9]+$ ]]; then 
    echo "This build was triggered by a Pull Request (PR)"
    exit 1 
fi 

# Checks if the build was triggered Pull Request.
#if [ "$TRAVIS_PULL_REQUEST" = "false" ] && [ ! "$TRAVIS_PULL_REQUEST_BRANCH" = "master" ]; then
#    echo "Run in a Pull Request"
#    exit 1
#fi

if [ ! "$TRAVIS_BRANCH" == "master" ] && [ -z "$TRAVIS_PULL_REQUEST_BRANCH" ]; then
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