package base

import (
	log "base/internal"
	"base/internal/micro"
	"base/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"os"
)
/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 11:36
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func InitServerBase(pconf string) (*BaseProject, error) {
	if len(pconf) == 0 {
		pconf = pathToConfig
	}

	base := &BaseProject{}

	err := base.initConfig(pconf)

	if err != nil {
		 return nil, err
	}

	base.initLogHandle()

	base.initMicro()

	err = base.initRouters()

	if err != nil {
		return nil, err
	}

	return base, nil
}

func InitServerBaseDev(pconf string) (*BaseProject, error) {
	if len(pconf) == 0 {
		pconf = pathToConfig
	}

	base := &BaseProject{}

	err := base.initConfig(pconf)
	base.initLogHandle()

	if err != nil {
		return nil, err
	}

	base.initMicro()

	err = base.initRouters()

	if err != nil {
		return nil, err
	}

	return base, nil
}

func (base *BaseProject) initMicro()  {
	base.gift = micro.InitGift(base.conf.ApiGiftDev, base.conf.ApiGiftProd)
}

func (base *BaseProject) initConfig(conf string) error {
	base.conf = &Config{}
	return utils.LoadConfig(conf, base.conf)
}

func (base *BaseProject) initDatabase() error  {
	var err error

	return err
}

func (base *BaseProject) initRouters() error {
	var err error
	if base.conf.ModeDebug == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	base.router = gin.New()
	base.router.Use(log.LoggerWithWriter(base.createFileLogRequests()))
	base.router.Use(gin.Recovery())


	base.router.ForwardedByClientIP = true
	base.router.Use(favicon.New(pathFavicon))
	base.router.GET("/", base.home)
	base.router.GET("/healthy", base.healthy)
	base.router.GET("/api/v2/gift", base.handleGiftV2)
	base.router.GET("/api/v2/dev/gift", base.handleGiftV2)

	//apiProd2 := base.router.Group("/api/v2")
	//apiProd2.GET("/gift", base.handleGiftV2)

	return err
}

func (base *BaseProject) ListenAndServe() {
	if base.conf.ModeDebug == 0 {
		fmt.Printf("Listening and serving HTTP on %s\n", base.conf.ServerAddr)
	}
	base.host, _ = os.Hostname()

	err := base.router.Run(base.conf.ServerAddr)
	if err != nil {
		panic(err)
	}
}