package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	resp, err := cli.Grant(context.Background(), 5)
	if err != nil {
		fmt.Printf("create lease failed, %v\n", err)
		return
	}
	//cancel()
	// 5秒钟之后, foo 这个key就会被移除
	_, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(clientv3.LeaseID(resp.ID)))
	if err != nil {
		fmt.Printf("put  etcd failed, %v\n", err)
		return
	}
	time.Sleep(time.Second * 4)
	r, err := cli.Get(context.TODO(), "foo")
	if err != nil {
		fmt.Printf("get  etcd failed, %v\n", err)
		return
	}
	for _, kv := range r.Kvs {
		fmt.Printf("%s:%s\n", kv.Key, kv.Value)
	}
}
