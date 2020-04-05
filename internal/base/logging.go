package base

import (
	"github.com/natefinch/lumberjack"
	"log"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 05/04/2020 22:24
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (base *BaseProject) initLogHandle() {
	//TODO: Tạo file logging
	outputFile := &lumberjack.Logger{
		Filename:   base.conf.LogHandler,
		MaxSize:    128, // megabytes
		MaxBackups: 2,
		MaxAge:     7, //days
	}
	log.SetOutput(outputFile)
	//// Để có thể đóng file một cách an toàn, tạo 1 goroutine chạy và lắng nghe các tín hiệu của OS được chỉ định
	//sigs := make(chan os.Signal, 1)
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//go func() {
	//	<-sigs
	//	utils.HandlePrintf("Service shutdown!!!")
	//	err := outputFile.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
}

func (base *BaseProject) createFileLogRequests() *lumberjack.Logger {
	// File log request (Theo cấu trúc sẵn có của GIN)
	logRequestFile := &lumberjack.Logger{
		Filename:   base.conf.LogRequest,
		MaxSize:    128, // megabytes
		MaxBackups: 2,
		MaxAge:     7, //days
	}

	//// Để có thể đóng file một cách an toàn, tạo 1 goroutine chạy và lắng nghe các tín hiệu của OS được chỉ định
	//sigs := make(chan os.Signal, 1)
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//
	//go func() {
	//	<-sigs
	//	err := logRequestFile.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//}()
	return logRequestFile
}