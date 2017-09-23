package main

import (
	"asuka/conf"
	"asuka/http"
)

func main() {
	conf.Init()
	go http.Server()
}
