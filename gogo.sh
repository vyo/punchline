#!/bin/sh

echo "setting up punchline workspace"
export GOPATH=$(pwd)

go get gopkg.in/alecthomas/kingpin.v2

