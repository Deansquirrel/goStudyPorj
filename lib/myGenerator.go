package lib

import (
	"context"
	"time"
)

const (
	STATUS_ORIGINAL uint32 = 0
	STATUS_STARTING uint32 = 1
	STATUS_STARTED  uint32 = 2
	STATUS_STOPING  uint32 = 3
	STATUS_STOPED   uint32 = 4
)

type Generator struct {
	caller      ICaller       //调用器
	timeoutNS   time.Duration //超时时间
	lps         uint32        //每秒载荷量
	durationNS  time.Duration //负载持续时间
	concurrency uint32        //载荷并发量

	tickets IGoTickets //goroutine票池

	ctx       context.Context    //上下文
	cancelFun context.CancelFunc //取消函数

	callCount uint64 //调用计数

	status uint32 //状态

	resultCh chan *CallResult //调用结果通道
}

func NewGenerator(caller ICaller, timeoutNS time.Duration, lps uint32, durationNS time.Duration, resultCh chan *CallResult) IGenerator {
	gen := &Generator{
		caller:     caller,
		timeoutNS:  timeoutNS,
		lps:        lps,
		durationNS: durationNS,
		status:     STATUS_ORIGINAL,
		resultCh:   resultCh,
	}
	return gen
}

func (g *Generator) Start() bool {
	return false
}

func (g *Generator) Stop() bool {
	return false
}

func (g *Generator) Status() uint32 {
	return g.status
}

func (g *Generator) CallCount() uint64 {
	return g.callCount
}
