#!/usr/bin/env sh

# NAME: install.sh
# DECRIPTION: Install helper script
# AUTHOR: Tony Cavella (tony@cavella.com)
# SOURCE: https://github.com/acavella/gorevoke

set -e

# Set build variables
GOREVOKE_TAG="v1.0.0-rc12"
APP_NAME="gorevoke"
APP_VERSION=${GOREVOKE_TAG#v}
APP_BUILD="Dev"
APP_BUILDTIME=$(date +%Y%m%d-%M%H)
BUILD_DIR="/usr/local/bin/gorevoke"
GOOS=linux
GOARCH=amd64

# Clone repository
git clone https://github.com/acavella/gorevoke
cd gorevoke
git checkout ${GOREVOKE_TAG}

# Create application base directories
mkdir -p ${BUILD_DIR}/conf
mkdir -p ${BUILD_DIR}/crl/tmp
mkdir -p ${BUILD_DIR}/crl/static

# Copy default files
cp ./conf/config.yml ${BUILD_DIR}/conf/ 

# Build Go application
go build -o "${BUILD_DIR}/${APP_NAME}" -ldflags="-X 'main.appVersion=${APP_VERSION}' -X 'main.appBuild=${APP_BUILD}' -X 'main.appBuildDate=${APP_BUILDTIME}'" main.go
