#!/bin/bash
# Works with a semver file VERSION in current directory,
# Display current version, automatically suggest a "patch" 
# version update, ask for input to use suggestion, 
# or a newly entered value and create a GIT tag.

if [ ! -f VERSION ]; then
    echo "0.0.0" > VERSION
fi

BASE_STRING=`cat VERSION`
BASE_LIST=(`echo $BASE_STRING | tr '.' ' '`)
V_MAJOR=${BASE_LIST[0]}
V_MINOR=${BASE_LIST[1]}
V_PATCH=${BASE_LIST[2]}
echo "Current version : $BASE_STRING"
V_PATCH=$((V_PATCH + 1))
SUGGESTED_VERSION="$V_MAJOR.$V_MINOR.$V_PATCH"
read -p "Enter a version number [$SUGGESTED_VERSION]: " INPUT_STRING
if [ "$INPUT_STRING" = "" ]; then
    INPUT_STRING=$SUGGESTED_VERSION
fi
echo "Will set new version to be $INPUT_STRING"
echo $INPUT_STRING > VERSION
git add VERSION
git commit -m "$INPUT_STRING"
git tag -a -m "Tagging version $INPUT_STRING" "v$INPUT_STRING"
