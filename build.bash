#!/usr/bin/env bash
set -ex

# This script builds archiver for most common platforms.

# https://github.com/uber/h3-go/issues/23
export CGO_ENABLED=0

cd cmd/
# https://www.sakishum.com/2021/11/29/Golang-%E4%BA%A4%E5%8F%89%E7%BC%96%E8%AF%91%E6%8A%A5%E9%94%99-XX-is-invalid-in-C99/
GOOS=linux   GOARCH=amd64 go build -o ../builds/arc_linux_amd64
#GOOS=linux   GOARCH=amd64 CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ go build -o ../builds/arc_linux_amd64
# https://www.anycodings.com/1questions/91017/cross-compiling-for-linux-arm7-clang-error-argument-unused-during-compilation-marm
# GOOS=linux   GOARCH=arm   go build -o ../builds/arc_linux_arm7
GOOS=darwin  GOARCH=amd64 go build -o ../builds/arc_mac_amd64
# GOOS=windows GOARCH=amd64 go build -o ../builds/arc_windows_amd64.exe
cd ../
