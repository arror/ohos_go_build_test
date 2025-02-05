#!/bin/bash

OHOS_NATIVE_HOME="/home/runner/work/ohos_go_build_test/ohos_go_build_test/command-line-tools/sdk/default/openharmony/native"

export AR=$OHOS_NATIVE_HOME/llvm/bin/llvm-ar
export CC=$OHOS_NATIVE_HOME/llvm/bin/aarch64-unknown-linux-ohos-clang
export CXX=$OHOS_NATIVE_HOME/llvm/bin/aarch64-unknown-linux-ohos-clang++
export GOOS=openharmony
export GOARCH=arm64
export CGO_ENABLED=1

go build -tls=GD -buildmode c-shared -o ./output/libohtest.so ./
