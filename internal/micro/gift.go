package micro

import (
	"base/internal/genjson"
	"base/internal/utils"
	"context"
	"encoding/json"
	"github.com/graph-gophers/dataloader"
	"github.com/patrickmn/go-cache"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
	"strings"
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
	//utils.HandlePrintf("- GetOpenGiftFeature")
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
	var results []*dataloader.Result

	reqMap := make(map[string][]string, 0)

	// map: key = userID - value = []string{postID}
	for _, obj := range keys {
		tmp := strings.Split(obj.String(), ":")
		// nếu không có userID => bỏ qua
		if len(tmp) < 2 {
			utils.HandleWarnPrintf("GetOpenGiftFeatureProdFunc len(key) < 2. Not found userID or postID")
			continue
		}
		// tồn tại userID trong map => thêm tiếp postID vào
		if val, ok := reqMap[tmp[0]]; ok {
			postIDs := append(val, tmp[1])
			reqMap[tmp[0]] = postIDs
		} else { // chưa có userID trong map => thêm mới
			reqMap[tmp[0]] = []string{tmp[1]}
		}
	}

	responseMap := make(map[string]map[string]bool)

	// map: key userID - value = map[postID]bool
	//O(total_userID)
	for userID, postIDs := range reqMap {
		// với mỗi userID có listPostID, call api một lần
		// TODO cần sử dụng goroutine?
		responseMap[userID] = mg.GetOpenGiftFeature(context.TODO(), userID, false, postIDs)
	}

	// O(len(keys))
	for _, obj := range keys {
		key := strings.Split(obj.String(), ":")
		userID, postID := key[0], key[1]
		// O(1)
		// map: key userID - value = map[postID]bool
		if tmp, ok := responseMap[userID]; ok {
			if val, ok := tmp[postID]; ok {
				results = append(results, &dataloader.Result{Data: val})
			} else {
				results = append(results, &dataloader.Result{Data: 0})
			}
		} else {
			results = append(results, &dataloader.Result{Data: 0})
		}
	}
	return results
}

func (mg *MicroGift) GetGiftCacheDev(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result

	reqMap := make(map[string][]string, 0)

	// map: key = userID - value = []string{postID}
	for _, obj := range keys {
		tmp := strings.Split(obj.String(), ":")
		// nếu không có userID => bỏ qua
		if len(tmp) < 2 {
			utils.HandleWarnPrintf("getOpenGiftFeatureDevFunc len(key) < 2. Not found userID or postID")
			continue
		}
		// tồn tại userID trong map => thêm tiếp postID vào
		if val, ok := reqMap[tmp[0]]; ok {
			postIDs := append(val, tmp[1])
			reqMap[tmp[0]] = postIDs
		} else { // chưa có userID trong map => thêm mới
			reqMap[tmp[0]] = []string{tmp[1]}
		}
	}

	responseMap := make(map[string]map[string]bool)

	// map: key userID - value = map[postID]bool
	// O(total_userID)
	for userID, postIDs := range reqMap {
		// với mỗi userID có listPostID, call api một lần
		// TODO cần sử dụng goroutine?
		responseMap[userID] = mg.GetOpenGiftFeature(context.TODO(), userID, true, postIDs)
	}

	// O(len(keys))
	for _, obj := range keys {
		key := strings.Split(obj.String(), ":")
		userID, postID := key[0], key[1]
		// O(1)
		// map: key = userID - value = map[postID]bool
		if tmp, ok := responseMap[userID]; ok {
			if val, ok := tmp[postID]; ok {
				results = append(results, &dataloader.Result{Data: val})
			} else {
				results = append(results, &dataloader.Result{Data: 0})
			}
		} else {
			results = append(results, &dataloader.Result{Data: 0})
		}
	}
	return results
}
