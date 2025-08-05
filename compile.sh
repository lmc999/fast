#!/bin/sh

echo "开始编译..."

# 使用单个Docker命令编译所有版本
docker run --rm -v $PWD:/app -w /app golang:1.21-alpine sh -c '
    apk add --no-cache git
    
    # 下载依赖
    go mod download
    
    # 创建输出目录
    mkdir -p dist
    
    # Linux amd64
    echo "编译 Linux amd64..."
    GOOS=linux GOARCH=amd64 go build -o dist/fast-linux-amd64 fast.go
    
    # Linux arm64
    echo "编译 Linux arm64..."
    GOOS=linux GOARCH=arm64 go build -o dist/fast-linux-arm64 fast.go
    
    # Windows amd64
    echo "编译 Windows amd64..."
    GOOS=windows GOARCH=amd64 go build -o dist/fast-windows-amd64.exe fast.go
    
    # macOS amd64
    echo "编译 macOS amd64..."
    GOOS=darwin GOARCH=amd64 go build -o dist/fast-darwin-amd64 fast.go
    
    # macOS arm64
    echo "编译 macOS arm64..."
    GOOS=darwin GOARCH=arm64 go build -o dist/fast-darwin-arm64 fast.go
'

echo "编译完成!"
ls -la dist/