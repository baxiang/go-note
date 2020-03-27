package main

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
)

//type Host struct {
//	Name    string `json:"name"`
//	Address string `json:"address"`
//	Port    int    `json:"port"`
//}
type HostYML struct {
	Name    string `yml:"name"`
	Address string `yml:"address"`
	Port    int    `yml:"port"`
}
func main() {
	// 加载配置文件
	if err := config.Load(file.NewSource(
		file.WithPath("./config/config.yml"),
		file.WithPath("./config/config.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}
	var host HostYML
	if err := config.Get("hosts", "database").Scan(&host); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(host.Name, host.Address, host.Port)
}
