package server

import (
	"time"

	"github.com/wendy512/go-iecp5/asdu"
)

type CommandHandler interface {
	// OnInterrogation 总召唤请求
	OnInterrogation(asdu.Connect, *asdu.ASDU, asdu.QualifierOfInterrogation) error
	// OnCounterInterrogation 总计数器请求
	OnCounterInterrogation(asdu.Connect, *asdu.ASDU, asdu.QualifierCountCall) error
	// OnRead 读定值请求
	OnRead(asdu.Connect, *asdu.ASDU, asdu.InfoObjAddr) error
	// OnClockSync 时钟同步请求
	OnClockSync(asdu.Connect, *asdu.ASDU, time.Time) error
	// OnResetProcess 进程重置请求
	OnResetProcess(asdu.Connect, *asdu.ASDU, asdu.QualifierOfResetProcessCmd) error
	// OnDelayAcquisition 延迟获取请求
	OnDelayAcquisition(asdu.Connect, *asdu.ASDU, uint16) error
	// OnTestCommand 测试命令请求
	OnTestCommand(asdu.Connect, *asdu.ASDU) error
	// OnASDU 控制命令请求
	OnASDU(asdu.Connect, *asdu.ASDU) error
}
