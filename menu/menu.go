package menu

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"asuka/conf"
	h "asuka/http"
)

type res struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func Init() error {
	token, err := h.GetAccessToken()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(conf.Menu)
	if err != nil {
		return err
	}

	body := bytes.NewReader(data)
	url := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + token
	contentType := "application/x-www-form-urlencoded"

	r, err := http.Post(url, contentType, body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var v res
	err = json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	if v.Errcode != 0 {
		return errors.New(string(b))
	}

	return nil
}
