#!/bin/env bash

# Perform steps I added, then delegate to modified Makefile from official
# OpenWhisk Go example.

# Meant to run from developer machine right now, not from CI/CD. Login command
# will prompt. Requires ibmcloud CLI to already be logged in.

GCP_CREDS_BASE64=$1

REGION="us-east"
RESOURCE_GROUP="charter-challenge"
FUNCTIONS_NAMESPACE="charter-challenge"

ibmcloud target -r $REGION
ibmcloud target -g $RESOURCE_GROUP
ibmcloud fn namespace target $FUNCTIONS_NAMESPACE

make deploy GCP_CREDS_BASE64=${GCP_CREDS_BASE64}
