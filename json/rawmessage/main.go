package main

import ("encoding/json"
"fmt")

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func main() {
	data := `{"hello":"world"}`
	response := &Response{
		Code: 1,
		Msg:  "success",
		Data: json.RawMessage(data),
	}

	b, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(b))
}
