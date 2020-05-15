#/bin/bash

# The build script is tested on Fedora 32

BASE_NAME="FileGen"

PATH_SCRIPT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

PATH_BUILD="$PATH_SCRIPT/../bin"
mkdir -p $PATH_BUILD

cd "$PATH_SCRIPT/../"



OS=Linux

ARCH=64
GOOS=linux
GOARCH=amd64
echo "Compiling... (OS: $OS, Arch: $ARCH)"
GOOS=$GOOS GOARCH=$GOARCH go install github.com/gotk3/gotk3/gtk
GOOS=$GOOS GOARCH=$GOARCH go build -o "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}"

# ARCH=32
# GOOS=linux
# GOARCH=386
# echo "Compiling... (OS: $OS, Arch: $ARCH)"
# GOOS=$GOOS GOARCH=$GOARCH go install github.com/gotk3/gotk3/gtk
# GOOS=$GOOS GOARCH=$GOARCH go build -o "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}"



OS=Windows
TEMP_OUT="${PATH_BUILD}/${OS}-temp"

ARCH=64
GOOS=windows
GOARCH=amd64
echo "Compiling... (OS: $OS, Arch: $ARCH)"
rm -rf $TEMP_OUT
mkdir -p $TEMP_OUT
PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=$GOOS GOARCH=$GOARCH go install github.com/gotk3/gotk3/gtk
PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=$GOOS GOARCH=$GOARCH go build -o "${TEMP_OUT}/${BASE_NAME}.exe" -ldflags="-H windowsgui"
PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=$GOOS GOARCH=$GOARCH go build -o "${TEMP_OUT}/${BASE_NAME}-cli.exe"
cp -r /usr/x86_64-w64-mingw32/sys-root/mingw/{bin/*.dll,share} $TEMP_OUT
cd $TEMP_OUT && nsiswrapper --run ${BASE_NAME}.exe ${BASE_NAME}-cli.exe ./* --outfile "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}.exe"
rm -rf $TEMP_OUT

# ARCH=32
# GOOS=windows
# GOARCH=386
# echo "Compiling... (OS: $OS, Arch: $ARCH)"
# rm -rf $TEMP_OUT
# mkdir -p $TEMP_OUT
# PKG_CONFIG_PATH=/usr/i686-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=i686-w64-mingw32-gcc GOOS=$GOOS GOARCH=$GOARCH go install github.com/gotk3/gotk3/gtk
# PKG_CONFIG_PATH=/usr/i686-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=i686-w64-mingw32-gcc GOOS=$GOOS GOARCH=$GOARCH go build -o "${TEMP_OUT}/${BASE_NAME}.exe" -ldflags="-H windowsgui"
# PKG_CONFIG_PATH=/usr/i686-w64-mingw32/sys-root/mingw/lib/pkgconfig CGO_ENABLED=1 CC=i686-w64-mingw32-gcc GOOS=$GOOS GOARCH=$GOARCH go build -o "${TEMP_OUT}/${BASE_NAME}-cli.exe"
# cp -r /usr/i686-w64-mingw32/sys-root/mingw/{bin/*.dll,share} $TEMP_OUT
# cd $TEMP_OUT && nsiswrapper --run ${BASE_NAME}.exe ${BASE_NAME}-cli.exe ./* --outfile "${PATH_BUILD}/${OS}${ARCH}_${BASE_NAME}.exe"
# rm -rf $TEMP_OUT
