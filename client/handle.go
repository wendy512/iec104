package client

import "github.com/wendy512/go-iecp5/asdu"

const (
	SinglePoint                            DataType = iota // 单点信息
	DoublePoint                                            // 双点信息
	MeasuredValueScaled                                    // 测量值，标度化值信息
	MeasuredValueNormal                                    // 测量值,规一化值信息
	StepPosition                                           // 步位置信息
	BitString32                                            // 比特位串信息
	MeasuredValueFloat                                     // 测量值,短浮点数信息
	IntegratedTotals                                       // 累计量信息
	EventOfProtectionEquipment                             // 继电器保护设备事件信息
	PackedStartEventsOfProtectionEquipment                 // 继电器保护设备事件信息
	PackedOutputCircuitInfo                                // 继电器保护设备成组输出电路信息
	PackedSinglePointWithSCD                               // 带变位检出的成组单点信息
	SingleCommandInfo
	DoubleCommandInfo
	StepCommandInfo
	SetPointCommandNormalInfo
	SetPointCommandScaledInfo
	SetPointCommandFloatInfo
	BitsString32CommandInfo
	UNKNOWN // 未知的
)

type DataType int

type clientHandler struct {
	call ASDUCall
}

// InterrogationHandler 总召唤回复
func (h *clientHandler) InterrogationHandler(_ asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnInterrogation(rxAsdu)
}

// CounterInterrogationHandler 总计数器回复
func (h *clientHandler) CounterInterrogationHandler(_ asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnCounterInterrogation(rxAsdu)
}

// ReadHandler 读定值回复
func (h *clientHandler) ReadHandler(_ asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnRead(rxAsdu)
}

// TestCommandHandler 测试下发回复
func (h *clientHandler) TestCommandHandler(_ asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnTestCommand(rxAsdu)
}

// ClockSyncHandler 时钟同步回复
func (h *clientHandler) ClockSyncHandler(_ asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnClockSync(rxAsdu)
}

// ResetProcessHandler 进程重置回复
func (h *clientHandler) ResetProcessHandler(_ asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnResetProcess(rxAsdu)
}

// DelayAcquisitionHandler 延迟获取回复
func (h *clientHandler) DelayAcquisitionHandler(_ asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnDelayAcquisition(rxAsdu)
}

// ASDUHandler ASDU上报，ASDU数据
func (h *clientHandler) ASDUHandler(_ asdu.Connect, rxAsdu *asdu.ASDU) error {
	return h.call.OnASDU(rxAsdu)
}

func GetDataType(typeId asdu.TypeID) DataType {
	switch typeId {
	case asdu.M_SP_NA_1, asdu.M_SP_TA_1, asdu.M_SP_TB_1:
		return SinglePoint
	case asdu.M_DP_NA_1, asdu.M_DP_TA_1, asdu.M_DP_TB_1:
		return DoublePoint
	case asdu.M_ST_NA_1, asdu.M_ST_TA_1, asdu.M_ST_TB_1:
		return StepPosition
	case asdu.M_BO_NA_1, asdu.M_BO_TA_1, asdu.M_BO_TB_1:
		return BitString32
	case asdu.M_ME_NB_1, asdu.M_ME_TB_1, asdu.M_ME_TE_1:
		return MeasuredValueScaled
	case asdu.M_ME_NA_1, asdu.M_ME_TA_1, asdu.M_ME_TD_1, asdu.M_ME_ND_1:
		return MeasuredValueNormal
	case asdu.M_ME_NC_1, asdu.M_ME_TC_1, asdu.M_ME_TF_1:
		return MeasuredValueFloat
	case asdu.M_IT_NA_1, asdu.M_IT_TA_1, asdu.M_IT_TB_1:
		return IntegratedTotals
	case asdu.M_EP_TA_1, asdu.M_EP_TD_1:
		return EventOfProtectionEquipment
	case asdu.M_EP_TB_1, asdu.M_EP_TE_1:
		return PackedStartEventsOfProtectionEquipment
	case asdu.M_EP_TC_1, asdu.M_EP_TF_1:
		return PackedOutputCircuitInfo
	case asdu.M_PS_NA_1:
		return PackedSinglePointWithSCD
	case asdu.C_SC_NA_1, asdu.C_SC_TA_1:
		return SingleCommandInfo
	case asdu.C_DC_NA_1, asdu.C_DC_TA_1:
		return DoubleCommandInfo
	case asdu.C_RC_NA_1, asdu.C_RC_TA_1:
		return StepCommandInfo
	case asdu.C_SE_NA_1, asdu.C_SE_TA_1:
		return SetPointCommandNormalInfo
	case asdu.C_SE_NB_1, asdu.C_SE_TB_1:
		return SetPointCommandScaledInfo
	case asdu.C_SE_NC_1, asdu.C_SE_TC_1:
		return SetPointCommandFloatInfo
	case asdu.C_BO_NA_1, asdu.C_BO_TA_1:
		return BitsString32CommandInfo
	default:
		return UNKNOWN
	}
}
