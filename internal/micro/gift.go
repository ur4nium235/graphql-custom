package micro

import (
	"base/internal/genjson"
	"base/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/patrickmn/go-cache"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
	"time"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 05/04/2020 20:03
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func InitGift(apiDev, apiProd string) *MicroGift {
	utils.HandlePrintf("- Init Microservice Gift")
	var mg = &MicroGift{}
	mg.client = utils.CreateFastClient()
	mg.APIGiftProd = apiProd
	mg.APIGiftDev = apiDev
	cacheFake := &dataloader.NoCache{}
	_ = cacheFake

	mg.GiftLoaderDev = dataloader.NewBatchedLoader(mg.GetGiftCacheDev, dataloader.WithCache(&Cache{c: cache.New(10*time.Second, 10*time.Second)}), dataloader.WithBatchCapacity(50))
	mg.GiftLoaderProd = dataloader.NewBatchedLoader(mg.GetGiftCacheProd, dataloader.WithCache(&Cache{c: cache.New(10*time.Second, 10*time.Second)}), dataloader.WithBatchCapacity(50))
	//mg.GiftLoaderDev = dataloader.NewBatchedLoader(mg.GetGiftCacheDev, dataloader.WithCache(cacheFake), dataloader.WithBatchCapacity(50))
	//mg.GiftLoaderProd = dataloader.NewBatchedLoader(mg.GetGiftCacheProd, dataloader.WithCache(cacheFake), dataloader.WithBatchCapacity(50))
	return mg
}

func (mg *MicroGift) GetOpenGiftFeature(_ context.Context, userID string, isDev bool, listIds []string) map[string]bool {
	utils.HandlePrintf("- GetOpenGiftFeature")
	if len(userID) == 0 || len(listIds) == 0 {
		return make(map[string]bool)
	}
	if mg.client == nil {
		mg.client = utils.CreateFastClient()
	}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	if isDev {
		req.SetRequestURI(mg.APIGiftDev)
	} else {
		req.SetRequestURI(mg.APIGiftProd)
	}
	req.Header.SetMethod("GET")
	req.Header.Add("user-id", userID)

	query := req.URI().QueryArgs()
	postIDs := ``
	for _, id := range listIds {
		postIDs = postIDs + "," + id
	}
	postIDs = postIDs[1:]
	query.Add("postId", postIDs)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := mg.client.DoTimeout(req, resp, 1*time.Second)

	var out = make(map[string]bool)
	if err == nil {
		bodyBytes := resp.Body()
		tmp := genjson.ResponseGift{}
		if resp.StatusCode() == http.StatusOK {
			err = json.Unmarshal(bodyBytes, &tmp)
			if err == nil && len(tmp.Result) > 0 {
				for _, post := range tmp.Result {
					if post.Status == 1 {
						out[post.PostID] = true
					} else {
						out[post.PostID] = false
					}
				}
			} else {
				e, _ := url.QueryUnescape(string(query.QueryString()))
				utils.HandleErrorPrintf(err, req.URI().String()+"?"+e)
			}
		}
	} else {
		e, _ := url.QueryUnescape(string(query.QueryString()))
		utils.HandleErrorPrintf(err, req.URI().String()+"?"+e)
	}

	utils.HandlePrintfMsg(out, "GetOpenGiftFeature")
	return out
}

func (mg *MicroGift) GetGiftCacheProd(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
	utils.HandlePrintf("GetGiftCacheProd")
	var results []*dataloader.Result

	if len(keys) < 2 {
		utils.HandleWarnPrintf("GetGiftCacheProd len(keys) < 2. Not found userID or listID")
	}

	userID := keys.Keys()[0]
	listID := keys.Keys()[1:]
	tmp := mg.GetOpenGiftFeature(context.TODO(), userID, false, listID)

	fmt.Println("GetGiftCacheProd", tmp)

	for _, obj := range listID {
		if _, ok := tmp[obj]; ok {
			results = append(results, &dataloader.Result{Data: tmp[obj]})
		} else {
			results = append(results, &dataloader.Result{Data: 0})
		}
	}
	return results
}

func (mg *MicroGift) GetGiftCacheDev(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
	utils.HandlePrintf("GetGiftCacheDev")
	var results []*dataloader.Result

	if len(keys) < 2 {
		utils.HandleWarnPrintf("GetGiftCacheDev len(keys) < 2. Not found userID or listID")
	}

	userID := keys.Keys()[0]
	listID := keys.Keys()[1:]
	resp := mg.GetOpenGiftFeature(context.TODO(), userID, false, listID)
	for _, obj := range keys {
		if obj.String() == userID {
			continue
		}
		if _, ok := resp[obj.String()]; ok {
			results = append(results, &dataloader.Result{Data: resp[obj.String()]})
		} else {
			results = append(results, &dataloader.Result{Data: 0})
		}
	}
	return results
}
