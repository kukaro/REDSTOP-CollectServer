rm -rf ~/go
export GOPATH="~/go"

go get -u github.com/labstack/echo/...
go get -u github.com/go-sql-driver/mysql
go get -u golang.org/x/net/websocket
go get -u github.com/googollee/go-socket.io
go get -u github.com/rs/cors