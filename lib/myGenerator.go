package lib

import (
	"context"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"sync/atomic"
	"time"
)

const (
	StatusOriginal uint32 = 0
	StatusStarting uint32 = 1
	StatusStarted  uint32 = 2
	StatusStoping  uint32 = 3
	StatusStoped   uint32 = 4
)

const (
	RetCodeSuccessRetCode     = 0    //成功
	RetCodeWarningCallTimeout = 1001 //调用超时警告
	RetCodeErrorCall          = 2001 //调用错误
	RetCodeErrorResponse      = 2002 //相应内容错误
	RetCodeErrorCalee         = 2003 //被调用方(被测试软件)内部错误
	RetCodeFatalCall          = 3001 //调用过程中发射国内了致命错误
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

type ParamSet struct {
	Caller     ICaller          //调用器
	TimeoutNS  time.Duration    //超时时间
	Lps        uint32           //每秒载荷量
	DurationNS time.Duration    //负载持续时间
	ResultCh   chan *CallResult //调用结果通道
}

func NewGenerator(pset ParamSet) IGenerator {
	gen := &Generator{
		caller:     pset.Caller,
		timeoutNS:  pset.TimeoutNS,
		lps:        pset.Lps,
		durationNS: pset.DurationNS,
		status:     StatusOriginal,
		resultCh:   pset.ResultCh,
	}
	return gen
}

func (g *Generator) Start() bool {
	var throttle <-chan time.Time
	if g.lps > 0 {
		interval := time.Duration(1e9 / g.lps)
		fmt.Printf("Setting throttle (%v)...", interval)
		throttle = time.Tick(interval)
	} else {
		return false
	}
	g.ctx, g.cancelFun = context.WithTimeout(context.Background(), g.durationNS)
	g.callCount = 0
	atomic.StoreUint32(&g.status, StatusStarted)

	return true
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

func (g *Generator) genLoad(throttle <-chan time.Time) {
	for {
		//select {
		//	case <-g.ctx.Done():
		//		g.prepareToStop(g.ctx.Err())
		//		return
		//	default:
		//}
		g.asyncCall()
		if g.lps > 0 {
			select {
			case <-throttle:
			case <-g.ctx.Done():
				g.prepareToStop(g.ctx.Err())
				return
			}
		}
	}
}

func (g *Generator) prepareToStop(ctxError error) {
	fmt.Printf("Prepare to stop load (cause:%s)...", ctxError)
	atomic.CompareAndSwapUint32(&g.status, StatusStarted, StatusStoping)
	fmt.Println("Closing result channel...")
	close(g.resultCh)
	atomic.StoreUint32(&g.status, StatusStoped)
}

func (g *Generator) asyncCall() {
	g.tickets.Take()
	go func() {
		defer func() {
			g.tickets.Return()
		}()
	}()

	rawReq := g.caller.BuildReq()

	//0-未调用或调用中,1-调用完成,2-调用超时
	var callStatus uint32
	timer := time.AfterFunc(g.timeoutNS, func() {
		if !atomic.CompareAndSwapUint32(&callStatus, 0, 2) {
			return
		}
		result := CallResult{
			ID:     rawReq.ID,
			Req:    rawReq,
			Code:   RetCodeWarningCallTimeout,
			Msg:    fmt.Sprintf("Timeout!(expected:< %v", g.timeoutNS),
			Elapse: g.timeoutNS,
		}
		g.sendResult(&result)
	})

}

func (g *Generator) sendResult(result *CallResult) bool {
	if atomic.LoadUint32(&g.status) != StatusStarted {
		return false
	}
	select {
	case g.resultCh <- result:
		return true
	default:
		return false
	}
}

func (g *Generator) callOne(rawReq *RawReq) *RawResp {
	atomic.AddUint64(&g.callCount, 1)
	if rawReq == nil {
		return &RawResp{
			ID:  -1,
			Err: errors.New("Invalid raw request"),
		}
	}
	start := time.Now().UnixNano()
	resp, err := g.caller.Call(rawReq.Req, g.timeoutNS)
	end := time.Now().UnixNano()
	elapsedTime := time.Duration(end - start)
	if err != nil {

	}
}
