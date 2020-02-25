package base

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 11:23
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (base *BaseProject) home(c *gin.Context) {
	c.Header("Content-Type", "application/octet-stream")
	c.String(http.StatusOK, "Hello world, this is Sync-Cache-KingHub")
}

func (base *BaseProject) healthy(c *gin.Context) {
	val := rand.Intn(500) + 200
	time.Sleep(time.Duration(val) * time.Millisecond)
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With")
	c.Header("Connection", "Keep-Alive")
	winNoticeImg, _ := hex.DecodeString("47494638396101000100800000" +
		"FFFFFF0000002C000000000100010000" +
		"02024401003B")
	c.Header("Content-Type", "image/gif")
	_, _ = c.Writer.Write(winNoticeImg)
}