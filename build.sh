#!/bin/bash

# 创建输出目录
mkdir -p dist

echo "开始编译多平台二进制文件..."

# Linux AMD64
echo "编译 Linux x86_64..."
docker run --rm -v "$PWD":/go/src/fast -w /go/src/fast \
    -e GOOS=linux -e GOARCH=amd64 \
    golang:alpine go build -ldflags="-s -w" -o dist/fast_linux_amd64

# Linux ARM64
echo "编译 Linux ARM64..."
docker run --rm -v "$PWD":/go/src/fast -w /go/src/fast \
    -e GOOS=linux -e GOARCH=arm64 \
    golang:alpine go build -ldflags="-s -w" -o dist/fast_linux_arm64

# macOS AMD64
echo "编译 macOS x86_64..."
docker run --rm -v "$PWD":/go/src/fast -w /go/src/fast \
    -e GOOS=darwin -e GOARCH=amd64 \
    golang:alpine go build -ldflags="-s -w" -o dist/fast_darwin_amd64

# macOS ARM64 (Apple Silicon)
echo "编译 macOS ARM64..."
docker run --rm -v "$PWD":/go/src/fast -w /go/src/fast \
    -e GOOS=darwin -e GOARCH=arm64 \
    golang:alpine go build -ldflags="-s -w" -o dist/fast_darwin_arm64

# Windows AMD64
echo "编译 Windows x86_64..."
docker run --rm -v "$PWD":/go/src/fast -w /go/src/fast \
    -e GOOS=windows -e GOARCH=amd64 \
    golang:alpine go build -ldflags="-s -w" -o dist/fast_windows_amd64.exe

# Windows ARM64
echo "编译 Windows ARM64..."
docker run --rm -v "$PWD":/go/src/fast -w /go/src/fast \
    -e GOOS=windows -e GOARCH=arm64 \
    golang:alpine go build -ldflags="-s -w" -o dist/fast_windows_arm64.exe

echo "编译完成！"
ls -lah dist/