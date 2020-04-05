package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"runtime"
	"time"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 09:16
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func LoadConfig(pathToFile string, config interface{}) error {
	_, err := os.Stat(pathToFile)
	if err != nil {
		return err
	}

	if _, err := toml.DecodeFile(pathToFile, config); err != nil {
		return err
	}
	return nil
}

func HandlePrintf(msg interface{}) {
	log.Printf("[I] %v", msg)
}

func HandlePrintfMsg(err interface{}, msg string) {
	log.Printf("[I] %v, %v", msg, err)
}

func HandleErrorPrintf(err interface{}, msg string) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		if len(msg) > 0 {
			log.Printf("[E] %v %s:%d \t %v", err, fn, line, msg)
		} else {
			log.Printf("[E] %v %s:%d", err, fn, line)
		}
	}
}

func HandleWarnPrintf(msg interface{}) {
	log.Printf("[W] %v", msg)
}

func HandleErrorFatalf(err interface{}) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Fatalf("[E] %v %s:%d", err, fn, line)
	}
}
func GenHash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func CreateFastClient() *fasthttp.Client {
	return &fasthttp.Client{MaxConnsPerHost: 20000,
		MaxIdleConnDuration: 10 * time.Second,
	}
}
