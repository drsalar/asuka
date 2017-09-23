package http

import (
	"strconv"

	"asuka/conf"
)

func Server() {
	wx_port := ":" + strconv.Itoa(conf.Http_wx.Port)
	http.HandleFunc("/wx", wx)
	err := http.ListenAndServeTLS(wx_port, conf.Liscensedir+"/server.crt", conf.Liscensedir+"/server.key", nil)
	if err != nil {
		panic(err.Error())
	}
}
