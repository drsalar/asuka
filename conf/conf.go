package conf

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	V string `json:"version"`
	H Http   `json:"http"`
	D string `json:"liscensedir"`
	R string `json:"runEnv"`
	L string `json:"logLevel"`
	M string `json:"menu"`
}

type Http struct {
	Wx      HttpInstance `json:"wx"`
	Wxapi   HttpInstance `json:"wxapi"`
	Wxtoken HttpInstance `json:"wxtoken"`
}

type HttpInstance struct {
	Port string `json:"port"`
}

var Http_wx, Http_wxapi, Http_wxtoken HttpInstance
var Version, Liscensedir, RunEnv, LogLevel, Menu string

func Init() error {
	file, err := ioutil.ReadFile("conf.json")
	if err != nil {
		return err
	}
	var c config
	err = json.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	Http_wx = c.H.Wx
	Http_wxapi = c.H.Wxapi
	Http_wxtoken = c.H.Wxtoken
	Version = c.V
	Liscensedir = c.D
	RunEnv = c.R
	LogLevel = c.L
	Menu = c.M
	return nil
}
