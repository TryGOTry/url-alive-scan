@echo off
echo "build linux x64 admin & agent..."
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w" -o release/Web-Server-Scan E:\title-scan\main.go

echo "build windows x64 admin & agent..."
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w" -o release/Web-Server-Scan_x64.exe E:\title-scan\main.go