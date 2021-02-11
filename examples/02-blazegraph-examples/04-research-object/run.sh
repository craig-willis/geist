#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE WITH RO WITH BASE URI" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb
blazegraph import --file ../data/ro-with-base-uri.jsonld --format jsonld

END_SCRIPT

bash ${RUNNER} S1 "EXPORT RO WITH BASE URI" << END_SCRIPT

blazegraph export 

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE WITH RO NO BASE URI" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb
blazegraph import --file ../data/ro-no-base.jsonld --format jsonld

END_SCRIPT


bash ${RUNNER} S2 "EXPORT RO NO BASE URI" << END_SCRIPT

blazegraph export 

END_SCRIPT

# *****************************************************************************

bash ${RUNNER} SETUP "INITIALIZE BLAZEGRAPH INSTANCE WITH RO NO REMOTE CONTEXT" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb
blazegraph import --file ../data/ro-no-remote-context.jsonld --format jsonld

END_SCRIPT


bash ${RUNNER} S3 "EXPORT RO NO REMOTE CONTEXT" << END_SCRIPT

blazegraph export 

END_SCRIPT

# *****************************************************************************
