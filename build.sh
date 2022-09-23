#/bin/bash

REREASE="0.3.5"

git tag -a v$REREASE -m "version $REREASE"
git push origin v$REREASE
GOOS=windows GOARCH=amd64 go build -o ./build/elevator-simulator-$REREASE-amd64-windows/
GOOS=darwin GOARCH=arm64 go build -o ./build/elevator-simulator-$REREASE-arm64-darwin/
GOOS=linux GOARCH=amd64 go build -o ./build/elevator-simulator-$REREASE-amd64-linux/
GOOS=darwin GOARCH=amd64 go build -o ./build/elevator-simulator-$REREASE-amd64-darwin/
GOOS=linux GOARCH=arm64 go build -o ./build/elevator-simulator-$REREASE-arm64-linux/
cd ./build
zip ./elevator-simulator-$REREASE-arm64-darwin.zip -r ./elevator-simulator-$REREASE-arm64-darwin/
zip ./elevator-simulator-$REREASE-amd64-darwin.zip -r ./elevator-simulator-$REREASE-amd64-darwin/
zip ./elevator-simulator-$REREASE-arm64-linux.zip -r ./elevator-simulator-$REREASE-arm64-linux/
zip ./elevator-simulator-$REREASE-amd64-linux.zip -r ./elevator-simulator-$REREASE-amd64-linux/
zip ./elevator-simulator-$REREASE-amd64-windows.zip -r ./elevator-simulator-$REREASE-amd64-windows/
