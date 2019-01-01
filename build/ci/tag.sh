#!/bin/bash

if [ $TRAVIS ]; then

    if ! [ $TRAVIS_PULL_REQUEST ]; then

        if [ $TRAVIS_BRANCH == "master" ]; then

            echo $TRAVIS_BRANCH

            git config --global user.email "builds@travis-ci.com"
            git config --global user.name "Travis CI"
            echo `cat VERSION.txt`
            export GIT_TAG=`cat VERSION.txt`
            git tag $GIT_TAG -a -m "Generated tag from TravisCI build $TRAVIS_BUILD_NUMBER"
            echo `git show origin`
            git push -q https://$GITHUB_OAUTH_TOKEN@github.com/cameronnewman/go-emailvalidation --tag > /dev/null 2>&1

        fi

        echo "Build is not in the master branch"
        exit 1
    fi

    echo "Run in a Pull Request"
    exit 1
fi
echo "Script needs to be run on TravisCI platform"
exit 1