#!/bin/bash

set -eu

mkdir -p $GOPATH/src/github.com/engineerbetter/concourse-up
mv concourse-up/* $GOPATH/src/github.com/engineerbetter/concourse-up
cd $GOPATH/src/github.com/engineerbetter/concourse-up

deployment="system-test-$RANDOM"

go run main.go deploy $deployment

config=$(go run main.go info $deployment)

domain=$(echo $config | jq -r '.config.domain')
username=$(echo $config | jq -r '.config.concourse_username')
password=$(echo $config | jq -r '.config.concourse_password')

if [[ -n $TLS_CERT ]]; then
  fly --target system-test login --concourse-url https://$domain --username $username --password $password
else
  fly --target system-test login --insecure --concourse-url https://$domain --username $username --password $password
fi

set -x
fly --target system-test workers
set +x


# DESTROY

go run main.go --non-interactive destroy $deployment