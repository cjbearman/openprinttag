#!/bin/bash

# To ensure we can include the source_repo in the context, we build from parent directory
cd ..
docker build -t opttest:latest -f docker_test/Dockerfile .
