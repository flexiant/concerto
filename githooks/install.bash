#!/bin/bash
pwd=`git rev-parse --show-toplevel`
echo cp $pwd/githooks/pre-commit $pwd/.git/hooks/
