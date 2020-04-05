package genjson

/**
 *
 * @author: hoangtq
 * @timeCreate: 05/04/2020 16:35
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type ResponseGift struct {
	Result  []Gift `json:"result"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Gift struct {
	PostID  string `json:"postID"`
	Version string `json:"version"`
	IsSys   int    `json:"isSys"`
	Status  int    `json:"status"`
}

type ResponseClient struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Result  []Gift `json:"result"`
}
