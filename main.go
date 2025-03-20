package main

import (
	"C"
	"encoding/json"
	"vpn/log"
)
import (
	"net"
)

type Config struct {
	Fd   int32  `json:"fd"`
	Path string `json:"path"`
}

//export Startup
func Startup(config []byte) {
	cfg := Config{}
	json.Unmarshal(config, &cfg)
	if jsonStr, err := json.Marshal(cfg); err == nil {
		log.Debug("Config: " + string(jsonStr))
	} else {
		log.Error("Config: " + err.Error())
		return
	}

	// 使用不同的方式创建监听器
	addr := &net.TCPAddr{
		IP:   net.ParseIP("198.18.0.1"),
		Port: 0,
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Error("Listen failed: " + err.Error())
		return
	}

	log.Debug("Listening on: " + listener.Addr().String())
	listener.Close()
}

//export Shutdown
func Shutdown() {
	log.Debug("Shutdown")
}

func main() {}
