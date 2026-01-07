#!/bin/bash

go install golang.org/x/pkgsite/cmd/pkgsite@latest
echo ">> View docs at http://localhost:8080"
$(go env GOPATH)/bin/pkgsite
