package common

import "encoding/json"

type Status struct {
	Msg    string
	Code   int
	Result interface{}
}

func (s *Status) GetJson() []byte {
	j, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return j
}

func (s *Status) SetMsg(msg string) *Status {
	s.Msg = msg
	return s
}

func (s *Status) SetCode(code int) *Status {
	s.Code = code
	return s
}

func (s *Status) SetResult(result interface{}) *Status {
	s.Result = result
	return s
}

func GetSuccessStatus() *Status {
	return &Status{Msg:"", Code:200, Result:nil}
}

func GetErrorStatus() *Status {
	return &Status{Msg:"系统错误", Code:1000, Result:nil}
}

func GetCustomStatus(msg string, code int, result interface{}) *Status {
	return &Status{Msg:msg, Code:code, Result:result}
}
