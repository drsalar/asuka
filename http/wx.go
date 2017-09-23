package http

import (
	// "crypto/sha1"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	// "sort"
	// "strings"
	"time"

	// "asuka/conf"
	"asuka/log"
)

func wx(w http.ResponseWriter, r *http.Request) {
	// key := "ilovechina"
	// var s, t, n, echostr, data string
	// if r.Method == "GET" {
	// 	s = r.FormValue("signature")
	// 	t = r.FormValue("timestamp")
	// 	n = r.FormValue("nonce")
	// 	echostr = r.FormValue("echostr")
	// 	data = r.FormValue("data")
	// } else if r.Method == "POST" {
	// 	s = r.PostFormValue("signature")
	// 	t = r.PostFormValue("timestamp")
	// 	n = r.PostFormValue("nonce")
	// 	echostr = r.PostFormValue("echostr")
	// 	data = r.PostFormValue("data")
	// }

	// fmt.Println("****\n", r.Method, s, t, n, echostr, data, "****\n")

	// log.Debug("wx", "check signature", "", "signature", s, "timestamp", t, "nonce", n, "echostr", echostr, "data", data)

	// if conf.RunEnv == "online" {
	// 	ss := []string{t, n, key}
	// 	sort.Strings(ss)
	// 	sha := sha1.New()
	// 	io.WriteString(sha, strings.Join(ss, ""))
	// 	genSignature := fmt.Sprintf("%x", sha.Sum(nil))
	// 	if genSignature != s {
	// 		log.Warning("wx", "check signature", "signature not matched!", "signature", s, "genSignature", genSignature)
	// 		return
	// 	}
	// }

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("wx", "read data", "")
		return
	}

	data := string(body)
	log.Debug("wx", "read data", "", "data", data)

	res, err := dataHandler(data)
	if err != nil {
		log.Error("wx", "handle data", err.Error(), "data", data)
		return
	}

	log.Debug("wx", "return data", "", "res", res)

	w.Write([]byte(res))
	return
}

// <xml>
// <ToUserName><![CDATA[公众号]]></ToUserName>
// <FromUserName><![CDATA[粉丝号]]></FromUserName>
// <CreateTime>1460537339</CreateTime>
// <MsgType><![CDATA[text]]></MsgType>
// <Content><![CDATA[欢迎开启公众号开发者模式]]></Content>
// <MsgId>6272960105994287618</MsgId>
// </xml>

//  <xml>
// <ToUserName><![CDATA[粉丝号]]></ToUserName>
// <FromUserName><![CDATA[公众号]]></FromUserName>
// <CreateTime>1460541339</CreateTime>
// <MsgType><![CDATA[text]]></MsgType>
// <Content><![CDATA[test]]></Content>
// </xml>

type RetData struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CharData `xml:"ToUserName"`
	FromUserName CharData `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      CharData `xml:"MsgType"`
	Content      CharData `xml:"Content"`
}

type CharData struct {
	Value string `xml:",cdata"`
}

type RecData struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
}

func dataHandler(data string) (res string, err error) {
	var recData RecData
	var retData RetData
	err = xml.Unmarshal([]byte(data), &recData)
	if err != nil {
		return
	}

	// retData.Content = charData(recData.Content)
	// retData.ToUserName = charData(recData.FromUserName)
	// retData.FromUserName = charData(recData.ToUserName)
	// retData.MsgType = charData("text")
	// retData.CreateTime = charData(fmt.Sprintf("%v", time.Now().Unix()))
	retData.Content.Value = recData.Content
	retData.ToUserName.Value = recData.FromUserName
	retData.FromUserName.Value = recData.ToUserName
	retData.MsgType.Value = "text"
	retData.CreateTime = fmt.Sprintf("%v", time.Now().Unix())
	bytes, err := xml.Marshal(retData)
	if err != nil {
		return
	}

	log.Debug("dataHandler", "handler", "", "from", recData.FromUserName, "to", recData.ToUserName, "type", recData.MsgType, "createTime", recData.CreateTime, "content", recData.Content, "msgId", recData.MsgId)

	res = string(bytes)
	return
}

func charData(data string) string {
	return "<![CDATA[" + data + "]]>"
}
