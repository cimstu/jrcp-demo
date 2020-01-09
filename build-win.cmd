color 0a
cd /d "%~dp0\"
setlocal enabledelayedexpansion
title build..

set "PATH=%cd%\bin;%PATH%"
set "GOPATH=%cd%\gopath"
set "GOARCH=amd64"
set "CGO_ENABLED=1"

go build -o bin\client.exe client
go build -buildmode=c-shared -o bin\server.dll server
g++ cpp\main.cpp bin\server.dll -o bin\server.exe

start server.exe
start client.exe