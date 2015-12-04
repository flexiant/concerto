#!/bin/bash
#
pwd=`git rev-parse --show-toplevel`
cp $pwd/githooks/pre-commit $pwd/.git/hooks/
