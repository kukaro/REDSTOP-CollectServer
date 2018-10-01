package socket

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/websocket"
)

func Routers() *echo.Echo {
	e := echo.New()
	g := e.Group("/ws")
	{
		g.GET("/sign-in/:username/:password", getSignIn)
	}
	return e
}

func getSignIn(c echo.Context) error {
	username := c.Param("username")
	password := c.Param("password")
	fmt.Println(username, password)
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		//err := websocket.Message.Send(ws, "Hello, Client!")
		//if err != nil {
		//	c.Logger().Error(err)
		//}

		// Read
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			c.Logger().Error(err)
		}
		if len(msg) > 0 {
			fmt.Println("socket msg:" + msg)
			log.Debugf("socket msg:" + msg)
		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
