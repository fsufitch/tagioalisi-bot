#!/usr/bin/env bash

go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
    go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install -v github.com/google/wire/cmd/wire@latest
