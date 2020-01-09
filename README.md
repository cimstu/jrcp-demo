# jrcp-demo

## Build On Windows
运行脚本build-win.cmd会自动拉起rpc server、client

## Build On Mac
根目录下运行：

export GOPATH="$PWD/gopath"

export GOOS=darwin  GOARCH=amd64

export CGO_ENABLED=1

go build -o bin/client client

go build -buildmode=c-shared -o bin/server.so server

g++ cpp/main.cpp bin/server.so -o bin/server

二进制输出在bin目录下，终端窗口启动server、再启动client

## 依赖
golang

g++

## 输出
编译成功二进制输出在bin目录下：

可执行文件client --- rpc客户端（golang）

动态库server.dll(.so) --- rpc服务端（golang）

可执行文件server --- rpc服务端宿主（c++）

如编译有问题release目录下有已编译的可执行文件
