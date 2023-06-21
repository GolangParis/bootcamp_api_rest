#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

docker stop postgres-sandbox-instance 2>/dev/null
docker rm postgres-sandbox-instance 2>/dev/null

docker rmi postgres-sandbox 2>/dev/null

( cd ${SCRIPT_DIR} ; docker build -t postgres-sandbox -f Dockerfile.badges . )

docker run --name postgres-sandbox-instance -p 5432:5432 -i -t postgres-sandbox
