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

# Checks if the VERSION file exists from the makefile
if [ ! -f VERSION.txt ]; then
    echo "VERSION.txt file not found!"
    exit 2
fi

# Checks if the Github OAuth Token was set.
if [ -z "$GITHUB_OAUTH_TOKEN" ]; then
    echo "The GITHUB_OAUTH_TOKEN has not been set"
    exit 2
fi

export GIT_TAG=`cat VERSION.txt`

echo "Tagging release version $GIT_TAG"
echo $TRAVIS_BRANCH

git config --global user.email "builds@travis-ci.com"
git config --global user.name "Travis CI"

git tag $GIT_TAG -a -m "Generated tag from TravisCI build $TRAVIS_BUILD_NUMBER"
echo `git show origin`
git push -q https://$GITHUB_OAUTH_TOKEN@github.com/$TRAVIS_REPO_SLUG --tag > /dev/null 2>&1

echo "Pushed tag to repo"
exit 1