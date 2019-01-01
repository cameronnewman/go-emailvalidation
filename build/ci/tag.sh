#!/bin/bash
if [ "$TRAVIS_BRANCH" == "master" ]; then
  // do the deploy

  git config --global user.email "builds@travis-ci.com"
  git config --global user.name "Travis CI"
  echo `cat VERSION.txt`
  export GIT_TAG=`cat VERSION.txt`
  git tag $GIT_TAG -a -m "Generated tag from TravisCI build $TRAVIS_BUILD_NUMBER"
  echo `git show origin`
  git push origin $GIT_TAG

fi