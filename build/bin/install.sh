#!/usr/bin/evn sh

set -e

# Set build variables
APP_NAME="gorevoke"
APP_VERSION=${GOREVOKE_TAG}
APP_BUILD="Docker"
APP_BUILDTIME=$(date +%Y%m%d-%M%H)
GOOS=linux
GOARCH=amd64

# Build Go application
go build -o ${APP_NAME} main.go -ldflags="-X 'main.appVersion=${APP_VERSION}' -X 'main.appBuild=${APP_BUILD}' -X 'main.appBuildDate=${APP_BUILDTIME}'"
