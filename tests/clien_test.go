package tests

import (
	"fmt"
	"github.com/wendy512/go-iecp5/asdu"
	"github.com/wendy512/iec104/client"
	"sync"
	"testing"
	"time"
)

const (
	commonAddr = 1
)

func TestClient(t *testing.T) {
	settings := client.NewSettings()
	settings.Host = "192.168.33.12"
	settings.LogCfg = &client.LogCfg{Enable: true}
	c := client.New(settings, &clientCall{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	c.SetOnConnectHandler(func(c *client.Client) {
		// 连接成功以后做的操作
		fmt.Printf("connected %s iec104 server\n", settings.Host)
	})

	// server active确认后回调
	c.SetServerActiveHandler(func(c *client.Client) {
		// 发送总召唤
		if err := c.SendInterrogationCmd(commonAddr); err != nil {
			t.Errorf("send interrogation cmd error %v\n", err)
			t.FailNow()
		}

		// 累积量召唤
		if err := c.SendCounterInterrogationCmd(commonAddr); err != nil {
			t.Errorf("send counter interrogation cmd error %v\n", err)
			t.FailNow()
		}

		// read cmd
		if err := c.SendReadCmd(commonAddr, 400); err != nil {
			t.Errorf("send counter interrogation cmd error %v\n", err)
			t.FailNow()
		}

		// 时钟同步
		if err := c.SendClockSynchronizationCmd(commonAddr, time.Now()); err != nil {
			t.Errorf("send clock sync cmd error %v\n", err)
			t.FailNow()
		}

		// test cmd
		if err := c.SendTestCmd(commonAddr); err != nil {
			t.Errorf("send test cmd error %v\n", err)
			t.FailNow()
		}

		// 单点控制
		if err := c.SendCmd(commonAddr, asdu.C_SC_NA_1, asdu.InfoObjAddr(1000), true); err != nil {
			t.Errorf("send single cmd error %v\n", err)
			t.FailNow()
		}

		// 测试等待回复，不能结束太快
		time.Sleep(time.Second * 10)
		wg.Done()
	})

	// Connect后会发送server active
	if err := c.Connect(); err != nil {
		t.Errorf("client connect error %v\n", err)
		t.FailNow()
	}
	wg.Wait()

	if err := c.Close(); err != nil {
		t.Errorf("close error %v\n", err)
		t.FailNow()
	}
}

type clientCall struct {
}

// OnInterrogation 总召唤回复
func (c *clientCall) OnInterrogation(packet *asdu.ASDU) error {
	addr, value := packet.GetInterrogationCmd()
	fmt.Printf("interrogation reply, addr: %d, value: %d\n", addr, value)
	return nil
}

// OnCounterInterrogation 总计数器回复
func (c *clientCall) OnCounterInterrogation(packet *asdu.ASDU) error {
	addr, value := packet.GetCounterInterrogationCmd()
	fmt.Printf("counter interrogation reply, addr: %d, request: 0x%02X, rreeze: 0x%02X\n",
		addr, value.Request, value.Freeze)
	return nil
}

// OnRead 读定值回复
func (c *clientCall) OnRead(packet *asdu.ASDU) error {
	return c.OnASDU(packet)
}

// OnTestCommand 测试下发回复
func (c *clientCall) OnTestCommand(packet *asdu.ASDU) error {
	addr, value := packet.GetTestCommand()
	fmt.Printf("test cmd reply, addr: %d, value: %t\n", addr, value)
	return nil
}

// OnClockSync 时钟同步回复
func (c *clientCall) OnClockSync(packet *asdu.ASDU) error {
	addr, value := packet.GetClockSynchronizationCmd()
	fmt.Printf("clock sync reply, addr: %d, value: %d\n", addr, value.UnixMilli())
	return nil
}

// OnResetProcess 进程重置回复
func (c *clientCall) OnResetProcess(packet *asdu.ASDU) error {
	addr, value := packet.GetResetProcessCmd()
	fmt.Printf("reset process reply, addr: %d, value: 0x%02X\n", addr, value)
	return nil
}

// OnDelayAcquisition 延迟获取回复
func (c *clientCall) OnDelayAcquisition(packet *asdu.ASDU) error {
	addr, value := packet.GetDelayAcquireCommand()
	fmt.Printf("delay acquisition reply, addr: %d, value: %d\n", addr, value)
	return nil
}

// OnASDU 数据正体
func (c *clientCall) OnASDU(packet *asdu.ASDU) error {
	// 读取设备数据
	switch client.GetDataType(packet.Type) {
	case client.SinglePoint:
		c.onSinglePoint(packet)
	case client.DoublePoint:
		c.onDoublePoint(packet)
	case client.MeasuredValueScaled:
		c.onMeasuredValueScaled(packet)
	case client.MeasuredValueNormal:
		c.onMeasuredValueNormal(packet)
	case client.StepPosition:
		c.onStepPosition(packet)
	case client.BitString32:
		c.onBitString32(packet)
	case client.MeasuredValueFloat:
		c.onMeasuredValueFloat(packet)
	case client.IntegratedTotals:
		c.onIntegratedTotals(packet)
	case client.EventOfProtectionEquipment:
		c.onEventOfProtectionEquipment(packet)
	case client.PackedStartEventsOfProtectionEquipment:
		c.onPackedStartEventsOfProtectionEquipment(packet)
	case client.PackedOutputCircuitInfo:
		c.onPackedOutputCircuitInfo(packet)
	case client.PackedSinglePointWithSCD:
		c.onPackedSinglePointWithSCD(packet)
	default:
		return nil
	}

	return nil
}

func (c *clientCall) onSinglePoint(packet *asdu.ASDU) {
	// [M_SP_NA_1], [M_SP_TA_1] or [M_SP_TB_1] 获取单点信息信息体集合
	for _, p := range packet.GetSinglePoint() {
		fmt.Printf("single point, ioa: %d, value: %v\n", p.Ioa, p.Value)
	}
}

func (c *clientCall) onDoublePoint(packet *asdu.ASDU) {
	// [M_DP_NA_1], [M_DP_TA_1] or [M_DP_TB_1] 获得双点信息体集合
	for _, p := range packet.GetDoublePoint() {
		fmt.Printf("double point, ioa: %d, value: %v\n", p.Ioa, p.Value)
	}
}

func (c *clientCall) onMeasuredValueScaled(packet *asdu.ASDU) {
	// [M_ME_NB_1], [M_ME_TB_1] or [M_ME_TE_1] 获得测量值，标度化值信息体集合
	for _, p := range packet.GetMeasuredValueScaled() {
		fmt.Printf("measured value scaled, ioa: %d, value: %v\n", p.Ioa, p.Value)
	}
}

func (c *clientCall) onMeasuredValueNormal(packet *asdu.ASDU) {
	// [M_ME_NA_1], [M_ME_TA_1],[ M_ME_TD_1] or [M_ME_ND_1] 获得测量值,规一化值信息体集合
	for _, p := range packet.GetMeasuredValueNormal() {
		fmt.Printf("measured value normal, ioa: %d, value: %v\n", p.Ioa, p.Value)
	}
}

func (c *clientCall) onStepPosition(packet *asdu.ASDU) {
	// [M_ST_NA_1], [M_ST_TA_1] or [M_ST_TB_1] 获得步位置信息体集合
	for _, p := range packet.GetStepPosition() {
		// state：false: 设备未在瞬变状态 true： 设备处于瞬变状态
		fmt.Printf("step position, ioa: %d, state: %t, value: %d\n", p.Ioa, p.Value.HasTransient, p.Value.Val)
	}
}

func (c *clientCall) onBitString32(packet *asdu.ASDU) {
	// [M_ME_NC_1], [M_ME_TC_1] or [M_ME_TF_1].获得测量值,短浮点数信息体集合
	for _, p := range packet.GetMeasuredValueFloat() {
		fmt.Printf("bigtstring32, ioa: %d, value: %v\n", p.Ioa, p.Value)
	}
}

func (c *clientCall) onMeasuredValueFloat(packet *asdu.ASDU) {
	// [M_ME_NC_1], [M_ME_TC_1] or [M_ME_TF_1].获得测量值,短浮点数信息体集合
	for _, p := range packet.GetMeasuredValueFloat() {
		fmt.Printf("measured value float, ioa: %d, value: %v\n", p.Ioa, p.Value)
	}
}

func (c *clientCall) onIntegratedTotals(packet *asdu.ASDU) {
	// [M_IT_NA_1], [M_IT_TA_1] or [M_IT_TB_1]. 获得累计量信息体集合
	for _, p := range packet.GetIntegratedTotals() {
		fmt.Printf("integrated totals, ioa: %d, count: %d, SQ: 0x%02X, CY: %t, CA: %t, IV: %t\n",
			p.Ioa, p.Value.CounterReading, p.Value.SeqNumber, p.Value.HasCarry, p.Value.IsAdjusted, p.Value.IsInvalid)
	}
}

func (c *clientCall) onEventOfProtectionEquipment(packet *asdu.ASDU) {
	// [M_EP_TA_1] [M_EP_TD_1] 获取继电器保护设备事件信息体
	for _, p := range packet.GetEventOfProtectionEquipment() {
		fmt.Printf("event of protection equipment, ioa: %d, event: %d, qdp: %d, mesc: %d, time: %d\n",
			p.Ioa, p.Event, p.Qdp, p.Msec, p.Time.UnixMilli())
	}
}

func (c *clientCall) onPackedStartEventsOfProtectionEquipment(packet *asdu.ASDU) {
	// [M_EP_TB_1] [M_EP_TE_1] 获取继电器保护设备事件信息体
	p := packet.GetPackedStartEventsOfProtectionEquipment()
	fmt.Printf("packed start events of protection equipment, ioa: %d, event: %d, qdp: %d, mesc: %d, time: %d\n",
		p.Ioa, p.Event, p.Qdp, p.Msec, p.Time.UnixMilli())
}

func (c *clientCall) onPackedOutputCircuitInfo(packet *asdu.ASDU) {
	// [M_EP_TC_1] [M_EP_TF_1] 获取继电器保护设备成组输出电路信息信息体
	p := packet.GetPackedOutputCircuitInfo()
	fmt.Printf("packed Output circuit, ioa: %d, qci: %d, qdp: %d, mesc: %d, time: %d\n",
		p.Ioa, p.Oci, p.Qdp, p.Msec, p.Time.UnixMilli())
}

func (c *clientCall) onPackedSinglePointWithSCD(packet *asdu.ASDU) {
	// [M_PS_NA_1]. 获得带变位检出的成组单点信息
	for _, p := range packet.GetPackedSinglePointWithSCD() {
		fmt.Printf("packed single point with SCD, ioa: %d, scd: %d, qds: %d\n", p.Ioa, p.Scd, p.Qds)
	}
}
