#!/bin/bash

if [ ! -d  coverage ]; then
    mkdir coverage
fi


go test ./... -coverprofile=./coverage/cover.out;
go tool cover -html=./coverage/cover.out

