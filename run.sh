#!/bin/bash

if [ -z "$1" ]
then
  echo "no project specified"
  exit 1
fi

if [ ! -d ./build ]
then
  mkdir ./build
fi

go build -o "./build/$1" "./src/$1" && \
"./build/$1"