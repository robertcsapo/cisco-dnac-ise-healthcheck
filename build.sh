#!/usr/bin/env bash

package="cisco-dnac-ise-healthcheck"
output_folder="bin/"
platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm" "linux/arm64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package'-'$GOOS"-"$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+=".exe"
    fi

    env CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o $output_folder$output_name
    if [ $? -ne 0 ]; then
        echo "An error has occurred! Aborting the script execution..."
        exit 1
    fi
done
