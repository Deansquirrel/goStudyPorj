package lib

import "time"

type CallResult struct {
	ID     int64         //ID
	Req    RawReq        //原生请求
	Resp   RawResp       //原生响应
	Code   uint32        //相应代码
	Msg    string        //结果成因的简述
	Elapse time.Duration //耗时
}

type RawReq struct {
	ID  int64
	Req []byte
}

type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}
