package models

type ArrayNum struct {
	ValorArray [][]int `json:"valorArrayOutput"`
}
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type ArrayNumRotate struct {
	ValorArrayRotate [][]int `json:"valorArrayImput"`
}
