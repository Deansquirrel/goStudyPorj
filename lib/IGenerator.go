package lib

type IGenerator interface {
	Start() bool
	Stop() bool
	Status() uint32
	CallCount() uint64
}
