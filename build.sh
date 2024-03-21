#!/bin/bash

APP_NAME="ipmi-pid-controller"
OUTPUT_DIR="bin"

# Ensure the output directory exists
mkdir -p ${OUTPUT_DIR}

# OS/Arch combinations
PLATFORMS=("darwin/amd64" "darwin/arm64" "windows/amd64" "windows/arm64" "linux/amd64" "linux/arm64" "freebsd/amd64" "freebsd/arm64")

# Compile for each platform
for PLATFORM in "${PLATFORMS[@]}"; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    OUTPUT="${OUTPUT_DIR}/${APP_NAME}-${GOOS}-${GOARCH}"

    if [ $GOOS = "windows" ]; then
        OUTPUT+='.exe'
    fi

    echo "Building for $PLATFORM..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT ./main.go
done

echo "Compilation completed."
