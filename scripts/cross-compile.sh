#/bin/bash

BASE_NAME="filegen-from-excel"

PATH_SCRIPT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

PATH_BUILD="$PATH_SCRIPT/../bin"
mkdir -p $PATH_BUILD

cd "$PATH_SCRIPT/../"



OS="linux"

ARCH="32"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=linux GOARCH=386 go build -o "${PATH_BUILD}/${OS}/${BASE_NAME}_${ARCH}"

ARCH="64"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=linux GOARCH=amd64 go build -o "${PATH_BUILD}/${OS}/${BASE_NAME}_${ARCH}"



OS="windows"

ARCH="32"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=windows GOARCH=386 go build -o "${PATH_BUILD}/${OS}/${BASE_NAME}_${ARCH}.exe"

ARCH="64"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=windows GOARCH=amd64 go build -o "${PATH_BUILD}/${OS}/${BASE_NAME}_${ARCH}.exe"



OS="mac"

ARCH="32"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=darwin GOARCH=386 go build -o "${PATH_BUILD}/${OS}/${BASE_NAME}_${ARCH}"

ARCH="64"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=darwin GOARCH=amd64 go build -o "${PATH_BUILD}/${OS}/${BASE_NAME}_${ARCH}"
