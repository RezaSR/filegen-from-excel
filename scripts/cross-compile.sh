#/bin/bash

BASE_NAME="filegen-from-excel"

PATH_SCRIPT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

PATH_BUILD="$PATH_SCRIPT/../bin"
mkdir -p $PATH_BUILD

cd "$PATH_SCRIPT/../"



OS="Linux"

ARCH="32"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=linux GOARCH=386 go build -o "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}"

ARCH="64"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=linux GOARCH=amd64 go build -o "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}"



OS="Windows"

ARCH="32"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=windows GOARCH=386 go build -o "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}.exe"

ARCH="64"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=windows GOARCH=amd64 go build -o "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}.exe"



OS="Mac"

ARCH="32"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=darwin GOARCH=386 go build -o "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}"

ARCH="64"
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=darwin GOARCH=amd64 go build -o "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}"
