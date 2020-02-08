package main

import (
	"fmt"
	 "github.com/baxiang/go-note/rpc/user_pb"
	 "github.com/golang/protobuf/proto"
)

func main() {
	u := &user_pb.UserInfo{Id: 1001,Name:"ming",Password:"123456"}

	pbData,err := proto.Marshal(u)
	if err != nil {
		fmt.Println("Marshaling error: ", err)
	}

	 user :=&user_pb.UserInfo{}
	err = proto.Unmarshal(pbData, user)

	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	fmt.Println(user)
}
