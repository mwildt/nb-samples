#!/bin/bash

GO_BUILD="CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ."

echo "run $GO_BUILD"
eval $GO_BUILD