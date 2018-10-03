#!/usr/bin/env bash
fuser -k 2000/tcp

rm -rf ~/go

export GOPATH=/home/ubuntu/go

go get -u github.com/labstack/echo/...
go get -u github.com/go-sql-driver/mysql
go get -u golang.org/x/net/websocket
go get -u github.com/googollee/go-socket.io
go get -u github.com/rs/cors
go get -u github.com/BurntSushi/toml

cd ~/CollectServer/CollectServer
go build ~/CollectServer/CollectServer/server.go
./server &