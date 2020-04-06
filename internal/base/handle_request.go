package base

import (
	"base/internal/genjson"
	"base/internal/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"
	"net/http"
	"strings"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 05/04/2020 21:50
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func (base *BaseProject) handleGiftV2(c *gin.Context) {
	var response = genjson.ResponseClient{}

	if len(c.Request.URL.RawQuery) > 0 {
		params := c.Request.URL.Query()
		var userID string
		var postIDList = make([]string, 0)
		if val, ok := params["postids"]; ok && len(val) > 0 && len(val[0]) > 0 {
			tmp := strings.Split(val[0], ",")
			postIDList = append(postIDList, tmp...)
		}

		userID = c.Request.Header.Get("user-id")

		keyList := make([]dataloader.Key, len(postIDList))
		for index := range postIDList {
			keyList[index] = dataloader.StringKey(userID + ":" + postIDList[index])
		}

		var ress []interface{}

		gifts := make(map[string]genjson.Gift)

		ress, _ = base.gift.GiftLoaderProd.LoadMany(context.TODO(), keyList)()

		//utils.HandlePrintfMsg(ress, "handle_request 46")

		if len(ress) > 0 {
			for index := range ress {
				res := ress[index]
				if val, ok := res.(bool); ok {
					postID := postIDList[index]
					status := 1
					if !val {
						status = 0
					}
					gifts[postID] = genjson.Gift{
						PostID: postID,
						Status: status,
					}
				}
			}
		}
		//utils.HandlePrintfMsg(gifts, "handle_request Prod 63")
		listGift := make([]genjson.Gift, 0)

		for _, val := range gifts  {
			listGift = append(listGift, val)
		}


		response.Result = listGift
		response.Status = 1
		response.Code = 200
		response.Message = "OK"
	} else {
		response.Result = make([]genjson.Gift, 0)
		response.Status = 0
		response.Code = 200
		response.Message = "Fail"
	}
	base.response(c, response)
}

func (base *BaseProject) handleGiftV2Dev(c *gin.Context) {
	var response = genjson.ResponseClient{}

	if len(c.Request.URL.RawQuery) > 0 {
		params := c.Request.URL.Query()
		var userID string
		var postIDList = make([]string, 0)
		if val, ok := params["postids"]; ok && len(val) > 0 && len(val[0]) > 0 {
			tmp := strings.Split(val[0], ",")
			postIDList = append(postIDList, tmp...)
		}

		userID = c.Request.Header.Get("user-id")

		keyList := make([]dataloader.Key, len(postIDList))
		for index := range postIDList {
			keyList[index] = dataloader.StringKey(userID + ":" + postIDList[index])
		}

		var ress []interface{}

		gifts := make(map[string]genjson.Gift)

		ress, _ = base.gift.GiftLoaderDev.LoadMany(context.TODO(), keyList)()

		utils.HandlePrintfMsg(ress, "handle_request Dev 109")

		if len(ress) > 0 {
			for index := range ress {
				res := ress[index]
				if val, ok := res.(bool); ok {
					postID := postIDList[index]
					status := 1
					if !val {
						status = 0
					}
					gifts[postID] = genjson.Gift{
						PostID: postID,
						Status: status,
					}
				}
			}
		}
		utils.HandlePrintfMsg(gifts, "handle_request Dev 127")
		listGift := make([]genjson.Gift, 0)

		for _, val := range gifts  {
			listGift = append(listGift, val)
		}


		response.Result = listGift
		response.Status = 1
		response.Code = 200
		response.Message = "OK"
	} else {
		response.Result = make([]genjson.Gift, 0)
		response.Status = 0
		response.Code = 200
		response.Message = "Fail"
	}
	base.response(c, response)
}

func (king *BaseProject) response(c *gin.Context, response genjson.ResponseClient) {
	defer func() {
		if err := recover(); err != nil {
			utils.HandleErrorPrintf(err, "")
		}
	}()
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, X-Requested-With")
	c.Header("Connection", "Keep-Alive")
	c.Header("KeepAliveTimeout", "5")
	data, err := response.MarshalJSON()
	if err != nil {
		c.String(http.StatusNoContent, "")
	} else {
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, string(data))
	}
}
