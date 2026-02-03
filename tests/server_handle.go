package tests

import (
	"fmt"
	"time"

	"github.com/wendy512/go-iecp5/asdu"
)

const (
	commonAddr = 1
)

type myServerHandler struct {
}

func (ms *myServerHandler) OnInterrogation(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierOfInterrogation) error {
	// 1. Single Point
	_ = asdu.Single(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.SinglePointInfo{
		Ioa:   100,
		Value: true,
		Qds:   asdu.QDSGood,
	})

	// 1.1 Single Point CP24
	_ = asdu.SingleCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.SinglePointInfo{
		Ioa:   101,
		Value: true,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 1.2 Single Point CP56
	_ = asdu.SingleCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.SinglePointInfo{
		Ioa:   102,
		Value: true,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 2. Double Point
	_ = asdu.Double(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.DoublePointInfo{
		Ioa:   200,
		Value: asdu.DPIDeterminedOn,
		Qds:   asdu.QDSGood,
	})

	// 2.1 Double Point CP24
	_ = asdu.DoubleCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.DoublePointInfo{
		Ioa:   201,
		Value: asdu.DPIDeterminedOn,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 2.2 Double Point CP56
	_ = asdu.DoubleCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.DoublePointInfo{
		Ioa:   202,
		Value: asdu.DPIDeterminedOn,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 3. Step Position
	_ = asdu.Step(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.StepPositionInfo{
		Ioa: 300,
		Value: asdu.StepPosition{
			Val:          10,
			HasTransient: false,
		},
		Qds: asdu.QDSGood,
	})

	// 3.1 Step Position CP24
	_ = asdu.StepCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.StepPositionInfo{
		Ioa: 301,
		Value: asdu.StepPosition{
			Val:          10,
			HasTransient: false,
		},
		Qds:  asdu.QDSGood,
		Time: time.Now(),
	})

	// 3.2 Step Position CP56
	_ = asdu.StepCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.StepPositionInfo{
		Ioa: 302,
		Value: asdu.StepPosition{
			Val:          10,
			HasTransient: false,
		},
		Qds:  asdu.QDSGood,
		Time: time.Now(),
	})

	// 4. Bit String 32
	_ = asdu.BitString32(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.BitString32Info{
		Ioa:   400,
		Value: 0x12345678,
		Qds:   asdu.QDSGood,
	})

	// 4.1 Bit String 32 CP24
	_ = asdu.BitString32CP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.BitString32Info{
		Ioa:   401,
		Value: 0x12345678,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 4.2 Bit String 32 CP56
	_ = asdu.BitString32CP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.BitString32Info{
		Ioa:   402,
		Value: 0x12345678,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 5. Measured Value Normal
	_ = asdu.MeasuredValueNormal(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.MeasuredValueNormalInfo{
		Ioa:   500,
		Value: asdu.Normalize(16384), // 0.5 * 32768
		Qds:   asdu.QDSGood,
	})

	// 5.1 Measured Value Normal CP24
	_ = asdu.MeasuredValueNormalCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.MeasuredValueNormalInfo{
		Ioa:   501,
		Value: asdu.Normalize(16384),
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 5.2 Measured Value Normal CP56
	_ = asdu.MeasuredValueNormalCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.MeasuredValueNormalInfo{
		Ioa:   502,
		Value: asdu.Normalize(16384),
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 5.3 Measured Value Normal No Quality
	_ = asdu.MeasuredValueNormalNoQuality(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.MeasuredValueNormalInfo{
		Ioa:   503,
		Value: asdu.Normalize(16384),
	})

	// 6. Measured Value Scaled
	_ = asdu.MeasuredValueScaled(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.MeasuredValueScaledInfo{
		Ioa:   600,
		Value: 100,
		Qds:   asdu.QDSGood,
	})

	// 6.1 Measured Value Scaled CP24
	_ = asdu.MeasuredValueScaledCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.MeasuredValueScaledInfo{
		Ioa:   601,
		Value: 100,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 6.2 Measured Value Scaled CP56
	_ = asdu.MeasuredValueScaledCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.MeasuredValueScaledInfo{
		Ioa:   602,
		Value: 100,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 7. Measured Value Float
	_ = asdu.MeasuredValueFloat(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.MeasuredValueFloatInfo{
		Ioa:   700,
		Value: 123.456,
		Qds:   asdu.QDSGood,
	})

	// 7.1 Measured Value Float CP24
	_ = asdu.MeasuredValueFloatCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.MeasuredValueFloatInfo{
		Ioa:   701,
		Value: 123.456,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 7.2 Measured Value Float CP56
	_ = asdu.MeasuredValueFloatCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.MeasuredValueFloatInfo{
		Ioa:   702,
		Value: 123.456,
		Qds:   asdu.QDSGood,
		Time:  time.Now(),
	})

	// 8. Integrated Totals
	_ = asdu.IntegratedTotals(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.BinaryCounterReadingInfo{
		Ioa: 800,
		Value: asdu.BinaryCounterReading{
			CounterReading: 999,
			SeqNumber:      1,
		},
	})

	// 8.1 Integrated Totals CP24
	_ = asdu.IntegratedTotalsCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.BinaryCounterReadingInfo{
		Ioa: 801,
		Value: asdu.BinaryCounterReading{
			CounterReading: 999,
			SeqNumber:      1,
		},
		Time: time.Now(),
	})

	// 8.2 Integrated Totals CP56
	_ = asdu.IntegratedTotalsCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.BinaryCounterReadingInfo{
		Ioa: 802,
		Value: asdu.BinaryCounterReading{
			CounterReading: 999,
			SeqNumber:      1,
		},
		Time: time.Now(),
	})

	// 9. Packed Single Point with SCD
	_ = asdu.PackedSinglePointWithSCD(conn, false, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.PackedSinglePointWithSCDInfo{
		Ioa: 900,
		Scd: 0x01,
		Qds: asdu.QDSGood,
	})

	// 10. Protection Events
	_ = asdu.EventOfProtectionEquipmentCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.EventOfProtectionEquipmentInfo{
		Ioa:   1000,
		Event: asdu.SEDeterminedOff,
		Msec:  500,
		Qdp:   asdu.QDPGood,
		Time:  time.Now(),
	})

	_ = asdu.EventOfProtectionEquipmentCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.EventOfProtectionEquipmentInfo{
		Ioa:   1001,
		Event: asdu.SEDeterminedOff,
		Msec:  500,
		Qdp:   asdu.QDPGood,
		Time:  time.Now(),
	})

	// 11. Packed Start Events
	_ = asdu.PackedStartEventsOfProtectionEquipmentCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.PackedStartEventsOfProtectionEquipmentInfo{
		Ioa:   1100,
		Event: asdu.SEPGeneralStart,
		Qdp:   asdu.QDPGood,
		Msec:  100,
		Time:  time.Now(),
	})

	_ = asdu.PackedStartEventsOfProtectionEquipmentCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.PackedStartEventsOfProtectionEquipmentInfo{
		Ioa:   1101,
		Event: asdu.SEPGeneralStart,
		Qdp:   asdu.QDPGood,
		Msec:  100,
		Time:  time.Now(),
	})

	// 12. Packed Output Circuit Info
	_ = asdu.PackedOutputCircuitInfoCP24Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.PackedOutputCircuitInfoInfo{
		Ioa:  1200,
		Oci:  asdu.OCIGeneralCommand,
		Qdp:  asdu.QDPGood,
		Msec: 200,
		Time: time.Now(),
	})

	_ = asdu.PackedOutputCircuitInfoCP56Time2a(conn, asdu.CauseOfTransmission{Cause: asdu.Spontaneous}, commonAddr, asdu.PackedOutputCircuitInfoInfo{
		Ioa:  1201,
		Oci:  asdu.OCIGeneralCommand,
		Qdp:  asdu.QDPGood,
		Msec: 200,
		Time: time.Now(),
	})

	return nil
}

func (ms *myServerHandler) OnCounterInterrogation(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierCountCall) error {
	//_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.CounterInterrogationCmd(conn, asdu.CauseOfTransmission{Cause: asdu.Activation}, commonAddr, asdu.QualifierCountCall{asdu.QCCGroup1, asdu.QCCFrzRead})
	//_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnRead(conn asdu.Connect, pack *asdu.ASDU, addr asdu.InfoObjAddr) error {
	//_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.Single(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.SinglePointInfo{
		Ioa:   addr,
		Value: true,
		Qds:   asdu.QDSGood,
	})
	//_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnClockSync(conn asdu.Connect, pack *asdu.ASDU, tm time.Time) error {
	//_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	now := time.Now()
	_ = asdu.ClockSynchronizationCmd(conn, asdu.CauseOfTransmission{Cause: asdu.Activation}, commonAddr, now)
	//_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnResetProcess(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierOfResetProcessCmd) error {
	//_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.ResetProcessCmd(conn, asdu.CauseOfTransmission{Cause: asdu.Activation}, commonAddr, asdu.QPRGeneralRest)
	//_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnDelayAcquisition(conn asdu.Connect, pack *asdu.ASDU, msec uint16) error {
	//_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.DelayAcquireCommand(conn, asdu.CauseOfTransmission{Cause: asdu.Activation}, commonAddr, msec)
	//_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnTestCommand(conn asdu.Connect, pack *asdu.ASDU) error {
	fmt.Printf("Received Test Command: Type=%d, Pack=%+v\n", pack.Type, pack)
	// Debug infoObj
	// Since infoObj is private, we can't access it directly easily without reflection or just trusting SendReplyMirror works if called correctly.
	// But we can check if GetTestCommand works.
	if pack.Type == asdu.C_TS_NA_1 {
		ioa, _ := pack.GetTestCommand()
		fmt.Printf("C_TS_NA_1 IOA: %d\n", ioa)
	}

	if err := pack.SendReplyMirror(conn, asdu.ActivationCon); err != nil {
		fmt.Printf("SendReplyMirror failed: %v\n", err)
		return err
	}
	return nil
}

func (ms *myServerHandler) OnASDU(conn asdu.Connect, pack *asdu.ASDU) error {
	if pack.Type == asdu.C_TS_TA_1 {
		return ms.OnTestCommand(conn, pack)
	}
	//_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	cmd := pack.GetSingleCmd()
	_ = asdu.SingleCmd(conn, pack.Type, pack.Coa, pack.CommonAddr, asdu.SingleCommandInfo{
		Ioa:   cmd.Ioa,
		Value: cmd.Value,
		Qoc:   cmd.Qoc,
	})
	//_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	return nil
}
