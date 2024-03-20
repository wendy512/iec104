package client

import "github.com/thinkgos/go-iecp5/asdu"

type clientHandler struct {
}

// InterrogationHandler 总召唤
func (h *clientHandler) InterrogationHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return nil
}

// CounterInterrogationHandler 总计数器
func (h *clientHandler) CounterInterrogationHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return nil
}

// ReadHandler 读定值
func (h *clientHandler) ReadHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return nil
}

// TestCommandHandler 测试下发
func (h *clientHandler) TestCommandHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return nil
}

// ClockSyncHandler 时钟同步
func (h *clientHandler) ClockSyncHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return nil
}

// ResetProcessHandler 进程重置
func (h *clientHandler) ResetProcessHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return nil
}

// DelayAcquisitionHandler 延迟获取
func (h *clientHandler) DelayAcquisitionHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return nil
}

// ASDUHandler ASDU上报
func (h *clientHandler) ASDUHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	return nil
}
