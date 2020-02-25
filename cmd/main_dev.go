package main

import (
	"base/internal/base"
	"runtime"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 09:00
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func runDev()  {
	//logrus.SetOutput(os.Stdout)
	//logrus.SetLevel(logrus.DebugLevel)
	runtime.GOMAXPROCS(runtime.NumCPU())
	server, err := base.InitServerBase("")
	if err != nil {
		panic(err)
	}
	server.ListenAndServe()
}

func main() {
	runDev()
}
