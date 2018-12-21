package lib

type IGoTickets interface {
	//获得一张票
	Take()
	//归还一张票
	Return()
	//票池是否已激活
	Active() bool
	//票的总数
	Total() uint32
	//剩余的票数
	Remainder() uint32
}
