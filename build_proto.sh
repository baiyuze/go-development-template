#!/bin/bash

protoc --go_out=. --go-grpc_out=. api/proto/hello.proto