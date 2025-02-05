#!/bin/bash

# 设置架构和目标
arch="arm64"  # 可选值：amd64, arm64
target="aarch64"  # 可选值：x86_64, aarch64
outdir="arm64-v8a"  # 可选值：x86_64, arm64-v8a

# 设置 OHOS_NATIVE_HOME
OHOS_NATIVE_HOME="downloads/command-line-tools/sdk/default/openharmony/native"

# 基础编译标志
BASE_FLAGS="-Wno-error --sysroot=$OHOS_NATIVE_HOME/sysroot "

# 工具链路径
TOOLCHAIN="$OHOS_NATIVE_HOME/llvm"

# 设置环境变量
export CC="$TOOLCHAIN/bin/clang"
export CXX="$TOOLCHAIN/bin/clang++"
export LD="$TOOLCHAIN/bin/clang"
export CGO_AR="$TOOLCHAIN/bin/llvm-ar"
export GOASM="$TOOLCHAIN/bin/llvm-as"
export GOOS="openharmony"
export GOARCH="$arch"
export GOARM=""
export CGO_ENABLED="1"
export CGO_CXXFLAGS=""
export CGO_CFLAGS="-Wno-error --target=$target-linux-ohos $BASE_FLAGS"
export CGO_LDFLAGS="-extld=$LD --sysroot=$OHOS_NATIVE_HOME/sysroot --target=$target-linux-ohos"

# 源文件和输出文件
sourceFile="./"
outputFile="./output/libohtest.so"

# 构建命令，生成共享库
go build -buildmode c-shared -tags "ohos" -gcflags="all=-N -l" -o $outputFile $sourceFile

# 检查编译结果
if [ -f "$outputFile" ]; then
    echo "success: $outputFile"
else
    echo "failed"
fi
