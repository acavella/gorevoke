#!/usr/bin/env sh

# NAME: install.sh
# DECRIPTION: Dockerfile install helper script
# AUTHOR: Tony Cavella (tony@cavella.com)
# SOURCE: https://github.com/acavella/gorevoke

set -e

# Set build variables
APP_NAME="gorevoke"
APP_VERSION=${1#v}
APP_BUILD="Docker"
APP_BUILDTIME=$(date +%Y%m%d-%M%H)
BUILD_DIR="/usr/local/bin/gorevoke"
GOOS=linux
GOARCH=amd64

# Create application base directories
mkdir ${BUILD_DIR}
#mkdir ${BUILD_DIR}/conf
#mkdir ${BUILD_DIR}/crl
#mkdir ${BUILD_DIR}/crl/tmp
#mkdir ${BUILD_DIR}/crl/static

# Copy default files
#cp ./conf/config.yml ${BUILD_DIR}/conf/ 

# Build Go application
go build -o "${BUILD_DIR}/${APP_NAME}" -ldflags="-X 'main.appVersion=${APP_VERSION}' -X 'main.appBuild=${APP_BUILD}' -X 'main.appBuildDate=${APP_BUILDTIME}'" *.go
