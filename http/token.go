package http

import (
	"io/ioutil"
	"net/http"

	"asuka/conf"
)

func GetAccessToken() (string, error) {
	url := "localhost" + conf.Http_wxtoken.Port + "/token"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
