#!/bin/bash

#
# ref:
# - https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications-es
# - https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
#

# go build -v -ldflags="-X 'main.Version=v1.0.0' -X 'github.com/romelgomez/golang-hello-mod/flags.User=$(id -u -n)' -X 'github.com/romelgomez/golang-hello-mod/flags.Time=$(date)'"

PROJECT="github.com/romelgomez/golang-hello-mod"
VERSION="-X 'main.Version=v1.0.0'"
USER="-X '$PROJECT/flags.User=$(id -u -n)'"
TIME="-X '$PROJECT/flags.Time=$(date)'"
DATE="-X $PROJECT/flags.Date=$(date +%Y-%m-%dT%H:%M:%S).000Z"

go build -v -ldflags="$VERSION $USER $TIME $DATE"
