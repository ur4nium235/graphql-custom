package micro

import (
	"github.com/graph-gophers/dataloader"
	"github.com/valyala/fasthttp"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 05/04/2020 20:03
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type MicroGift struct {
	client         *fasthttp.Client
	APIGiftProd    string
	APIGiftDev     string
	GiftLoaderProd *dataloader.Loader
	GiftLoaderDev  *dataloader.Loader
}
