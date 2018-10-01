package router

import (
	"../conf"
	"./socket"
	"fmt"
	"github.com/labstack/echo"
	"strings"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	hosts := make(map[string]*Host)

	hosts[conf.Conf.Server.DomainWebSocket] = &Host{socket.Routers()}
	return hosts
}

func RunSubDomains() {
	e := echo.New()
	hosts := InitRoutes()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		firstUri := strings.Split(req.RequestURI, "/")[1]
		fmt.Println(req.Host + "/" + firstUri)
		if host := hosts[req.Host+"/"+firstUri]; host == nil {
			/*error state*/
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}
		return
	})
	if err := e.Start(conf.Conf.Server.Addr); err != nil {
		fmt.Println(err)
	}
}
