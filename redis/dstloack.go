package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

func getLock(redisAddr, lockKey string, ex uint, retry int) error {
	if retry <= 0 {
		retry = 10
	}
	conn, err := redis.DialTimeout("tcp", redisAddr, time.Minute, time.Minute, time.Minute)
	if err != nil {
		fmt.Println("conn to redis failed, err:%v", err)
		return err
	}
	defer conn.Close()
	ts := time.Now() // as random value
	for i := 1; i <= retry; i++ {
		if i > 1 { // sleep if not first time
			time.Sleep(time.Second)
		}
		v, err := conn.Do("SET", lockKey, ts, "EX", retry, "NX")
		if err == nil {
			if v == nil {
				fmt.Println("get lock failed, retry times:", i)
			} else {
				fmt.Println("get lock success")
				break
			}
		} else {
			fmt.Println("get lock failed with err:", err)
		}
		if i >= retry {
			err = fmt.Errorf("get lock failed with max retry times.")
			return err
		}
	}
	return nil
}

func unLock(redisAddr, lockKey string) error {
	conn, err := redis.DialTimeout("tcp", redisAddr, time.Minute, time.Minute, time.Minute)
	if err != nil {
		fmt.Println("conn to redis failed, err:%v", err)
		return err
	}
	defer conn.Close()
	v, err := redis.Bool(conn.Do("DEL", lockKey))
	if err == nil {
		if v {
			fmt.Println("unLock success")
		} else {
			fmt.Println("unLock failed")
			return fmt.Errorf("unLock failed")
		}
	} else {
		fmt.Println("unLock failed, err:", err)
		return err
	}
	return nil
}

const (
	RedisAddr = "127.0.0.1:6379"
)

func main() {
	var wg sync.WaitGroup

	key := "dis_lock"

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			// getLock
			err := getLock(RedisAddr, key, 10, 10)
			if err != nil {
				fmt.Println(fmt.Sprintf("worker[%d] get lock failed:%v", id, err))
				return
			}
			// sleep for random
			for j := 0; j < 5; j++ {
				time.Sleep(time.Second)
				fmt.Println(fmt.Sprintf("worker[%d] hold lock for %ds", id, j+1))
			}
			// unLock
			err = unLock(RedisAddr, key)
			if err != nil {
				fmt.Println(fmt.Sprintf("worker[%d] unlock failed:%v", id, err))
			}
			fmt.Println(fmt.Sprintf("worker[%d] done", id))
		}(i)
	}

	wg.Wait()
	fmt.Println("demo is done!")
}
