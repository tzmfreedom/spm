#!/bin/sh

if [ $# != 1 ]; then
  echo "Usage: $0 [binary name]"
  exit 0
fi

GOOS=windows GOARCH=386 go build -o ./bin/$1-windows386.exe
GOOS=windows GOARCH=amd64 go build -o ./bin/$1-windows64.exe

GOOS=darwin GOARCH=386 go build -o ./bin/$1-darwin386
GOOS=darwin GOARCH=amd64 go build -o ./bin/$1-darwin64
