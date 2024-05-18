package dto

import "encoding/json"

type BaseResponse struct {
	Code uint16 `json:"code"` 
	Data json.RawMessage `json:"data"`
	Status string `json:"status"`
	Error string `json:"error"`
}