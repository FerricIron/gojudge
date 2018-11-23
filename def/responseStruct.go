package def

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	ErrCode   int    `json:"errCode"`
	JudgeNode int    `json:"judgeNode"`
	AllNode   int    `json:"allNode"`
	Msg       []byte `json:"msg"`
}

const (
	JudgeFinished = iota
	AcceptCode
	WrongAnwser
	ComplierError
	TimeLimitError
	ComlierTimeLimitError
	MemoryLimitError
	OuputLimitError
	RunTimeError
	OtherError = -1
)

func (resp *Response) StructToBytes() (data []byte, err error) {
	data, err = json.Marshal(resp)
	return
}

func (resp *Response) String() string {
	return fmt.Sprintf("ErrCode: %d current JudgeNode: %d AllNode: %d msg: %s", resp.ErrCode, resp.JudgeNode, resp.AllNode, resp.Msg)
}
