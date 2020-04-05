package base

import (
	"base/internal/micro"
	"github.com/aerospike/aerospike-client-go"
	"github.com/gin-gonic/gin"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 11:34
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type BaseProject struct {
	router *gin.Engine
	conf   *Config
	host   string

	//Cache
	aeroClient *aerospike.Client
	gift       *micro.MicroGift
}

type Config struct {
	ServerAddr string
	ModeDebug  int

	// cache aerospike
	AerospikeKingHubHosts string

	// my sql
	SqlKingHubHost string
	SqlKingHubName string
	ApiGiftProd    string
	ApiGiftDev     string
	LogRequest     string
	LogHandler     string
}
