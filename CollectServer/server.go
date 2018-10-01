package main

import (
	"fmt"
	"./conf"
	"./router"
)

func main() {
	if err := conf.Init(""); err == nil {
		fmt.Println("config success")
	}
	router.RunSubDomains()
}
