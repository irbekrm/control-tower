#!/bin/bash

# shellcheck disable=SC1091
source control-tower/ci/tasks/lib/test-setup.sh

# shellcheck disable=SC1091
source control-tower/ci/tasks/lib/letsencrypt.sh


handleVerboseMode

cp "$BINARY_PATH" ./cup
chmod +x ./cup

setDeploymentName crt

trapDefaultCleanup

echo "DEPLOY WITH LETSENCRYPT STAGING CERT, AND CUSTOM DOMAIN"

custom_domain="$deployment-auto-2.control-tower.engineerbetter.com"

if [ "$IAAS" = "GCP" ]
then
  custom_domain="$deployment-auto-2.gcp.engineerbetter.com"
fi

./cup deploy "$deployment" \
  --domain "$custom_domain"
sleep 60

config=$(./cup info --json "$deployment")
# shellcheck disable=SC2034
username=$(echo "$config" | jq -r '.config.concourse_username')
# shellcheck disable=SC2034
password=$(echo "$config" | jq -r '.config.concourse_password')
# shellcheck disable=SC2034
manifest="$(dirname "$0")/hello.yml"
# shellcheck disable=SC2034
job="hello"
# shellcheck disable=SC2034
domain="$custom_domain"

# Download the right version of fly from Concourse UI
updateFly "${domain}"

assertPipelineIsSettableAndRunnable
