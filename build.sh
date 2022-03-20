#!/bin/bash

GO_BUILD="CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ."

echo "run $GO_BUILD"
eval $GO_BUILD

branch=$(git rev-parse --abbrev-ref HEAD)
echo $branch

docker build -t $branch .
docker run -d $branch