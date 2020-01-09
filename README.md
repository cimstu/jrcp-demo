# jrcp-demo

## Build On Windows
运行脚本build-win.cmd会自动拉起rpc server、client

二进制输出在bin目录下

## Build On Mac
根目录下运行：

export GOPATH="$PWD/gopath"

export GOOS=darwin  GOARCH=amd64

export CGO_ENABLED=1

go build -o bin/client client

go build -buildmode=c-shared -o bin/server.so server

g++ cpp/main.cpp bin/server.so -o bin/server

二进制输出在bin目录下，先启动server再启动client

## 依赖
golang

g++

## Release
如编译有问题release目录下有编译好的可执行文件
