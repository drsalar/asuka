package http

import (
	"net/http"

	"asuka/conf"
)

func Server() {
	wx_port := conf.Http_wx.Port
	http.HandleFunc("/wx", wx)
	err := http.ListenAndServeTLS(wx_port, conf.Liscensedir+"/server.crt", conf.Liscensedir+"/server.key", nil)
	if err != nil {
		panic(err.Error())
	}
}
