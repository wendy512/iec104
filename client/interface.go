package client

import "github.com/wendy512/go-iecp5/asdu"

// ASDUCall  is the interface of client handler
type ASDUCall interface {
	// OnInterrogation 总召唤回复
	OnInterrogation(*asdu.ASDU) error
	// OnCounterInterrogation 总计数器回复
	OnCounterInterrogation(*asdu.ASDU) error
	// OnRead 读定值回复
	OnRead(*asdu.ASDU) error
	// OnTestCommand 测试下发回复
	OnTestCommand(*asdu.ASDU) error
	// OnClockSync 时钟同步回复
	OnClockSync(*asdu.ASDU) error
	// OnResetProcess 进程重置回复
	OnResetProcess(*asdu.ASDU) error
	// OnDelayAcquisition 延迟获取回复
	OnDelayAcquisition(*asdu.ASDU) error
	// OnASDU 数据回复或控制回复
	OnASDU(*asdu.ASDU) error
}
