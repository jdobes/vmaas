#!/usr/bin/bash

set -e

ORIGINAL_BRANCH="master"
VERSION_REGEXP="[0-9]+\.[0-9]+"
RELEASE_BRANCH_PREFIX="vmaas-"
DOCKER_COMPOSE_FILE="docker-compose.yml"
TRAVIS_FILE=".travis.yml"
PATCH_DOCKERFILES="*/Dockerfile */Dockerfile.rhel7"
ENV_FILES="conf/webapp.env conf/reposcan.env"

if [ "$(git rev-parse --abbrev-ref HEAD)" != "$ORIGINAL_BRANCH" ]; then
    echo "Please checkout $ORIGINAL_BRANCH branch to create new release."
    exit 1
fi

(
# Git root
cd "$(pwd)/$(git rev-parse --show-cdup)"
print_help=0
if [ "$1" == "-u" ]; then
    if [ "$2" != "" ]; then
        echo "Deleting tag: v$2"
        git tag -d "v$2"
        echo "Deleting branch: $RELEASE_BRANCH_PREFIX$2"
        git branch -D "$RELEASE_BRANCH_PREFIX$2"
    else
        print_help=1
    fi
elif [[ "$1" =~ $VERSION_REGEXP ]]; then
    echo "Creating new tag: v$1"
    git tag "v$1"

    echo "Creating new branch: $RELEASE_BRANCH_PREFIX$1"
    git checkout -b "$RELEASE_BRANCH_PREFIX$1"
    echo " Updating $DOCKER_COMPOSE_FILE..."
    sed -i "s/:latest/:$1/g" "$DOCKER_COMPOSE_FILE"
    echo " Updating $TRAVIS_FILE..."
    sed -i "s/- master/- $RELEASE_BRANCH_PREFIX$1/g" "$TRAVIS_FILE"
    for DOCKERFILE in $PATCH_DOCKERFILES; do
        echo " Updating $DOCKERFILE..."
        sed -i "s/ENV VMAAS_VERSION=latest/ENV VMAAS_VERSION=$1/g" "$DOCKERFILE"
        git add "$DOCKERFILE"
    done
    git add "$DOCKER_COMPOSE_FILE" "$TRAVIS_FILE"
    git commit -m "Update version to $1"
    git checkout "$ORIGINAL_BRANCH"
    echo ""
    echo "Run:"
    echo "    git push origin $RELEASE_BRANCH_PREFIX$1 && git push origin v$1"
else
    print_help=1
fi

if [ $print_help -eq 1 ]; then
    echo "This script creates a new release stuffs:"
    echo "- new tag 'vVERSION'"
    echo "- new branch 'vmaas-VERSION' from master updating:"
    echo "  - VMAAS_VERSION in dockerfiles"
    echo "  - image names in docker-compose.yml"
    echo "  - branch param in .travis.yml"
    echo ""
    echo "Usage:"
    echo "    $0 version"
    echo "    (version matching \"$VERSION_REGEXP\")"
    echo "Example:"
    echo "    $0 0.1"
    echo "Undo:"
    echo "    $0 -u 0.1"
fi
)
