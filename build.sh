#!/bin/zsh

#
# ref:
# - https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications-es
# - https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
#

PROJECT="github.com/romelgomez/golang-hello-mod/go_examples/flags"

VERSION="-X 'main.Version=v0.0.1'"

USER="-X '$PROJECT.User=$(id -u -n)'"

TIME="-X '$PROJECT.Time=$(date)'"

DATE="-X $PROJECT.Date=$(date +%Y-%m-%dT%H:%M:%S).000Z"

ENVIRONMENT="-X '$PROJECT.Environment=development'"

FLAGS="$VERSION $USER $TIME $DATE $ENVIRONMENT"

go build -v -ldflags="$FLAGS"
