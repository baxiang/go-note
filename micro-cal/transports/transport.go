package transports

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/baxiang/go-note/micro-cal/endpoints"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)



func DecodeCalRequest(_ context.Context,req *http.Request)(interface{},error){
	vars  :=mux.Vars(req)
	reqPara :=endpoints.CalRequest{}
	reqType ,ok:=vars["type"]
	if !ok{
		return nil,errors.New("请求参数错误")
	}
	reqPara.CalType = reqType

	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &reqPara)

	return reqPara, nil
}

func EncodeCalResponse(c context.Context,w http.ResponseWriter,response interface{})error{
	w.Header().Set("Content-type","application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}