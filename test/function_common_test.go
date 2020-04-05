package main

import (
	"base/internal/base"
	"base/internal/micro"
	"base/internal/utils"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"testing"
	"github.com/stretchr/testify/assert"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 09:27
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func TestGenHash(t *testing.T) {
	plainText := "vccorpadtechadmicrobigdataplatform"
	hash256 := "0faec427ba144f820c2f4faceb7393dd3cbe953b3dea48d53e011a0586b3fadd"
	assert.Equal(t, utils.GenHash(plainText), hash256)
}

func BenchmarkGenHash(b *testing.B)  {
	for test := 0; test < b.N; test++ {
		plainText := "vccorpadtechadmicrobigdataplatform"
		plainText += string(test)
		utils.GenHash(plainText)
	}
}

func TestGiftProd(t *testing.T)  {
	config := &base.Config{}
	err := utils.LoadConfig("configs/server.conf", config)

	if err != nil {
		utils.HandleErrorPrintf(err, "")
	}

	dataLoader := micro.InitGift(config.ApiGiftDev, config.ApiGiftProd)
	keyList := make([]dataloader.Key, 2)
	keyList[0] = dataloader.StringKey("123456789")
	keyList[0] = dataloader.StringKey("695927798413266944")
	fmt.Println(dataLoader.GetGiftCacheProd(nil, keyList))
}