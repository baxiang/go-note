package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
type wordQueryReq struct {
	Q string `json:"q"`
	Lang string `json:"lang"`
	Cmd []string `json:"cmd"`
}
type wordCommonReq struct {
	AppId string `json:"appId"`
	ClientId string `json:"clientId"`
	TraceId string `json:"traceId"`
}
func main() {
	reqB, _ := json.Marshal(wordQueryReq{Q:"今天北京天气",Lang:"zh",Cmd:[]string{"word","pos","ner","lemma","rank"}})
	reqC, _ := json.Marshal(wordCommonReq{AppId:"11",ClientId:"11",TraceId:"11"})
	value := url.Values{"b":{string(reqB)},"c":{string(reqC)}}
	fmt.Println(value)
	resp, err :=http.PostForm("nlu",value)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	b,err:= ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
