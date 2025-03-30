#!/bin/bash

GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./build/www main.go
# GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./build/www_arm main.go
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./build/www.exe main.go