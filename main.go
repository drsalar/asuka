package main

import (
	"asuka/conf"
	"asuka/http"
	"asuka/log"
	"asuka/menu"
	"fmt"
)

func main() {
	err := conf.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Init()

	err = menu.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	http.Server()
}
