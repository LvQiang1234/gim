package config

import (
	"gim/pkg/logger"

	"go.uber.org/zap"
)

func initProdConf() {
	RPCAddr = RPCAddrConf{
		ConnectRPCAddr:  "addrs:///127.0.0.1:50000",
		BusinessRPCAddr: "addrs:///127.0.0.1:50100",
		LogicRPCAddr:    "addrs:///127.0.0.1:50200",
	}

	Connect = ConnectConf{
		TCPListenAddr:        ":8080",
		WSListenAddr:         ":8081",
		RPCListenAddr:        ":50000",
		LocalAddr:            "127.0.0.1:50000",
		RedisIP:              "127.0.0.1:6379",
		RedisPassword:        "",
		PushRoomSubscribeNum: 100,
		PushAllSubscribeNum:  100,
	}

	Logic = LogicConf{
		MySQL:         "root:XX1516305754@tcp(127.0.0.1:3306)/gim?charset=utf8&parseTime=true",
		NSQIP:         "127.0.0.1:3306",
		RedisIP:       "127.0.0.1:6379",
		RedisPassword: "",
		RPCListenAddr: ":50100",
	}

	Business = BusinessConf{
		MySQL:         "root:XX1516305754@tcp(127.0.0.1:3306)/gim?charset=utf8&parseTime=true",
		NSQIP:         "127.0.0.1:3306",
		RedisIP:       "127.0.0.1:6379",
		RedisPassword: "",
		RPCListenAddr: ":50200",
	}

	logger.Level = zap.DebugLevel
	logger.Target = logger.Console
}
