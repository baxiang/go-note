package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"io"
	"log"
	"net/http"
)

func handleRedis(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
	var key =""
	key = query.Get("k")
	fmt.Println(key)
	if r.Method =="GET" {
		val, err := client.Get(key).Result()
		if err == redis.Nil {
			io.WriteString(w,"key does not exist")
		} else if err != nil {
			io.WriteString(w,err.Error())
		}else {
			io.WriteString(w,fmt.Sprintf("%s: %s\n",key,val))
		}
	}else if r.Method =="POST"{
		var val =""
		val = query.Get("v")
		err := client.Set(key, val, 0).Err()
		if err != nil {
			io.WriteString(w,err.Error())
		}
		io.WriteString(w,fmt.Sprintf("%s: %s\n",key,val))
	}
}

var client *redis.Client

func init()  {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(pong)
	}

}

func main() {
	http.HandleFunc("/",handleRedis)
	log.Fatal(http.ListenAndServe(":8090",nil))
}
