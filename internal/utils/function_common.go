package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"runtime"
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

func HandleError(err interface{})  {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[E] %v %s:%d", err, fn, line)
	}
}

func GenHash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}


