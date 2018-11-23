package def

import (
	"encoding/json"
)

type Response struct {
	ErrCode   int    `json:"errCode"`
	JudgeNode int    `json:"judgeNode"`
	AllNode   int    `json:"allNode"`
	Msg       []byte `json:"msg"`
}

const(
	JudgingResponseCode=iota
	AcceptCode
	WrongAnwser
	ComplierError
	TimeLimitError
	ComlierTimeLimitError
	MemoryLimitError
	OuputLimitError
	RunTimeError
	OtherError=-1
)

func (resp *Response)StructToBytes() (data []byte,err error){
	data,err =json.Marshal(resp)
	return
}