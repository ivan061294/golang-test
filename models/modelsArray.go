package models

type ArrayNum struct {
	ValorArray [][]int `json:"valorArray"`
}
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type ArrayNumRotate struct {
	ValorArrayRotate [][]int `json:"valorArrayRotate90"`
}
