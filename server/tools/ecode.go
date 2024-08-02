package tools

import "fmt"

type ECode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (e *ECode) String() string {
	return fmt.Sprintf("code:%d,message:%s", e.Code, e.Message)
}
