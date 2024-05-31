package client

import "github.com/thinkgos/go-iecp5/asdu"

type clientHandler struct {
	call ClientASDUCall
}

// InterrogationHandler 总召唤
func (h *clientHandler) InterrogationHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnInterrogation(rxAsdu)
}

// CounterInterrogationHandler 总计数器
func (h *clientHandler) CounterInterrogationHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnCounterInterrogation(rxAsdu)
}

// ReadHandler 读定值
func (h *clientHandler) ReadHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnRead(rxAsdu)
}

// TestCommandHandler 测试下发
func (h *clientHandler) TestCommandHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnTestCommand(rxAsdu)
}

// ClockSyncHandler 时钟同步
func (h *clientHandler) ClockSyncHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnClockSync(rxAsdu)
}

// ResetProcessHandler 进程重置
func (h *clientHandler) ResetProcessHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnResetProcess(rxAsdu)
}

// DelayAcquisitionHandler 延迟获取
func (h *clientHandler) DelayAcquisitionHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnDelayAcquisition(rxAsdu)
}

// ASDUHandler ASDU上报
func (h *clientHandler) ASDUHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnASDU(rxAsdu)
}
