package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect etcd failed,%v\n", err)
		return
	}
	defer cli.Close()
	_, err = cli.Put(context.TODO(), "hello", "etcd")
	if err != nil {
		fmt.Printf("put etcd failed, %v\n", err)
		return
	}
	resp, err := cli.Get(context.TODO(), "hello")
	if err != nil {
		fmt.Printf("get  etcd failed, %v\n", err)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("%s:%s\n", kv.Key, kv.Value)
	}
	_, err = cli.Delete(context.TODO(), "hello")
	if err != nil {
		fmt.Printf("get  etcd failed, %v\n", err)
		return
	}
}
