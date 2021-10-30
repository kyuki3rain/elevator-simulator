#/bin/bash

git tag -a v0.3.1 -m 'version 0.3.1'
GOOS=windows GOARCH=amd64 go build -o ./build/elevator-simulator-0.3.1-amd64-windows/
GOOS=darwin GOARCH=arm64 go build -o ./build/elevator-simulator-0.3.1-arm64-darwin/
GOOS=linux GOARCH=amd64 go build -o ./build/elevator-simulator-0.3.1-amd64-linux/
GOOS=darwin GOARCH=amd64 go build -o ./build/elevator-simulator-0.3.1-amd64-darwin/
GOOS=linux GOARCH=arm64 go build -o ./build/elevator-simulator-0.3.1-arm64-linux/
zip ./build/elevator-simulator-0.3.1-arm64-darwin.zip ./build/elevator-simulator-0.3.1-arm64-darwin/elevator-simulator
zip ./build/elevator-simulator-0.3.1-amd64-darwin.zip ./build/elevator-simulator-0.3.1-amd64-darwin/elevator-simulator
zip ./build/elevator-simulator-0.3.1-arm64-linux.zip ./build/elevator-simulator-0.3.1-arm64-linux/elevator-simulator
zip ./build/elevator-simulator-0.3.1-amd64-linux.zip ./build/elevator-simulator-0.3.1-amd64-linux/elevator-simulator
zip ./build/elevator-simulator-0.3.1-amd64-windows.zip ./build/elevator-simulator-0.3.1-amd64-windows/elevator-simulator.exe
