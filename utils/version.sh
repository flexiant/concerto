#!/bin/bash
release=`grep VERSION utils/version.go  | awk '{ print $4 }' | sed 's/\"//g'`
version=`echo $release | awk -F \. '{ print $1 }'`
mayor=`echo $release | awk -F \. '{ print $2 }'`
minor=`echo $release | awk -F \. '{ print $3 }'`
new_minor=$(( $minor + 1 ))
new_release="$version.$mayor.$new_minor"
echo "package utils"
echo ""
echo "const VERSION = \"$new_release\""