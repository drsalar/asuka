package main

import (
	"asuka/conf"
	"asuka/http"
	"asuka/log"
)

func main() {
	conf.Init()
	log.Init()
	http.Server()
}
