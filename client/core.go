package client

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cast"
	"github.com/wendy512/go-iecp5/asdu"
	"github.com/wendy512/go-iecp5/clog"
	"github.com/wendy512/go-iecp5/cs104"
	"github.com/wendy512/iec104/pkg/waitgroup"
	"net"
	"net/url"
	"strconv"
	"sync/atomic"
	"time"
)

type Client struct {
	client104             *cs104.Client
	settings              *Settings
	onConnectHandler      func(c *Client)
	connectionLostHandler func(c *Client)
}

// Settings 连接配置
type Settings struct {
	Host              string
	Port              int
	AutoConnect       bool          //自动重连
	ReconnectInterval time.Duration //重连间隔
	Cfg104            *cs104.Config //104协议规范配置
	TLS               *tls.Config   // tls配置
	Params            *asdu.Params  //ASDU相关特定参数
	LogCfg            *LogCfg
}

type LogCfg struct {
	Enable      bool //是否开启log
	LogProvider clog.LogProvider
}

type command struct {
	typeId asdu.TypeID
	ca     asdu.CommonAddr
	ioa    asdu.InfoObjAddr
	t      time.Time
	qoc    asdu.QualifierOfCommand
	qos    asdu.QualifierOfSetpointCmd
	value  any
}

func NewSettings() *Settings {
	cfg104 := cs104.DefaultConfig()
	return &Settings{
		Host:              "localhost",
		Port:              2404,
		AutoConnect:       true,
		ReconnectInterval: time.Minute,
		Cfg104:            &cfg104,
		Params:            asdu.ParamsWide,
	}
}

func New(settings *Settings, call ASDUCall) *Client {
	opts := newClientOption(settings)
	handler := &clientHandler{call: call}
	client104 := cs104.NewClient(handler, opts)
	logCfg := settings.LogCfg
	if logCfg != nil {
		client104.LogMode(logCfg.Enable)
		client104.SetLogProvider(logCfg.LogProvider)
	}

	return &Client{
		settings:  settings,
		client104: client104,
	}
}

func (c *Client) Connect() error {
	if err := c.testConnect(); err != nil {
		return err
	}

	if err := c.client104.Start(); err != nil {
		return err
	}

	wg := &waitgroup.WaitGroup{}
	wg.Add(1)
	// 标记是不是第一次
	var firstConnect atomic.Bool
	// 连接状态事件
	c.client104.SetOnConnectHandler(func(cs *cs104.Client) {
		if firstConnect.CompareAndSwap(false, true) {
			wg.Done()
		}
		cs.SendStartDt()
		if c.onConnectHandler != nil {
			c.onConnectHandler(c)
		}
	})

	// WaitTimeout 等待WaitGroup完成，但最多等待指定的持续时间
	if err := wg.WaitTimeout(c.settings.Cfg104.ConnectTimeout0); err != nil {
		return fmt.Errorf("connection timeout of %f seconds", c.settings.Cfg104.ConnectTimeout0.Seconds())
	}
	return nil
}

func (c *Client) Close() error {
	c.client104.SendStopDt()
	return c.client104.Close()
}

func (c *Client) SetLogCfg(cfg LogCfg) {
	c.client104.LogMode(cfg.Enable)
	c.client104.SetLogProvider(cfg.LogProvider)
}

func (c *Client) SetOnConnectHandler(f func(c *Client)) {
	c.onConnectHandler = f
}

func (c *Client) SetConnectionLostHandler(f func(c *Client)) {
	c.client104.SetConnectionLostHandler(func(_ *cs104.Client) {
		f(c)
	})
}

func (c *Client) SetServerActiveHandler(f func(c *Client)) {
	c.client104.SetServerActiveHandler(func(_ *cs104.Client) {
		f(c)
	})
}

func (c *Client) IsConnected() bool {
	return c.client104.IsConnected()
}

// SendCmd 双点遥控
func (c *Client) SendCmd(addr uint16, typeId asdu.TypeID, ioa asdu.InfoObjAddr, value any) error {
	cmd := &command{
		typeId: typeId,
		ioa:    ioa,
		ca:     asdu.CommonAddr(addr),
		value:  value,
		qoc: asdu.QualifierOfCommand{
			Qual:     asdu.QOCNoAdditionalDefinition,
			InSelect: false,
		},
		qos: asdu.QualifierOfSetpointCmd{
			Qual:     0,
			InSelect: false,
		},
		t: time.Now(),
	}

	return c.doSend(cmd)
}

// SendInterrogationCmd 发起总召唤
func (c *Client) SendInterrogationCmd(addr uint16) error {
	cmd := &command{typeId: asdu.C_IC_NA_1, ca: asdu.CommonAddr(addr)}
	return c.doSend(cmd)
}

// SendClockSynchronizationCmd 发起时钟同步
func (c *Client) SendClockSynchronizationCmd(addr uint16, t time.Time) error {
	cmd := &command{typeId: asdu.C_CS_NA_1, ca: asdu.CommonAddr(addr), t: t}
	return c.doSend(cmd)
}

// SendCounterInterrogationCmd 发起累积量召唤
func (c *Client) SendCounterInterrogationCmd(addr uint16) error {
	cmd := &command{typeId: asdu.C_CI_NA_1, ca: asdu.CommonAddr(addr)}
	return c.doSend(cmd)
}

// SendReadCmd 发起读命令
func (c *Client) SendReadCmd(addr uint16, ioa uint) error {
	cmd := &command{typeId: asdu.C_RD_NA_1, ioa: asdu.InfoObjAddr(ioa), ca: asdu.CommonAddr(addr)}
	return c.doSend(cmd)
}

// SendResetProcessCmd 发起复位进程命令
func (c *Client) SendResetProcessCmd(addr uint16) error {
	cmd := &command{typeId: asdu.C_RP_NA_1, ca: asdu.CommonAddr(addr)}
	return c.doSend(cmd)
}

// SendTestCmd 发送带时标的测试命令
func (c *Client) SendTestCmd(addr uint16) error {
	cmd := &command{typeId: asdu.C_TS_TA_1, ca: asdu.CommonAddr(addr)}
	return c.doSend(cmd)
}

func (c *Client) doSend(cmd *command) error {
	if !c.IsConnected() {
		return NotConnected
	}
	coa := activationCoa()
	var err error

	switch cmd.typeId {
	case asdu.C_IC_NA_1:
		err = c.client104.InterrogationCmd(coa, cmd.ca, asdu.QOIStation)
	case asdu.C_CI_NA_1:
		qcc := asdu.QualifierCountCall{Request: asdu.QCCTotal, Freeze: asdu.QCCFrzRead}
		err = c.client104.CounterInterrogationCmd(coa, cmd.ca, qcc)
	case asdu.C_CS_NA_1:
		err = c.client104.ClockSynchronizationCmd(coa, cmd.ca, cmd.t)
	case asdu.C_RD_NA_1:
		err = c.client104.ReadCmd(coa, cmd.ca, cmd.ioa)
	case asdu.C_RP_NA_1:
		err = c.client104.ResetProcessCmd(coa, cmd.ca, asdu.QPRGeneralRest)
	case asdu.C_TS_TA_1:
		err = c.client104.TestCommand(coa, cmd.ca)
	case asdu.C_SC_NA_1, asdu.C_SC_TA_1:
		var value bool
		value, err = cast.ToBoolE(cmd.value)
		if err != nil {
			return err
		}
		asduCmd := asdu.SingleCommandInfo{
			Ioa:   cmd.ioa,
			Value: value,
			Qoc:   cmd.qoc,
		}
		if cmd.typeId == asdu.C_SC_TA_1 {
			asduCmd.Time = cmd.t
		}
		err = asdu.SingleCmd(c.client104, cmd.typeId, coa, cmd.ca, asduCmd)
	case asdu.C_DC_NA_1, asdu.C_DC_TA_1:
		var value uint8
		value, err = cast.ToUint8E(cmd.value)
		if err != nil {
			return err
		}
		asduCmd := asdu.DoubleCommandInfo{
			Ioa:   cmd.ioa,
			Value: asdu.DoubleCommand(value),
			Qoc:   cmd.qoc,
		}
		if cmd.typeId == asdu.C_DC_TA_1 {
			asduCmd.Time = cmd.t
		}
		err = asdu.DoubleCmd(c.client104, cmd.typeId, coa, cmd.ca, asduCmd)
	case asdu.C_RC_NA_1, asdu.C_RC_TA_1:
		var value uint8
		value, err = cast.ToUint8E(cmd.value)
		if err != nil {
			return err
		}
		asduCmd := asdu.StepCommandInfo{
			Ioa:   cmd.ioa,
			Value: asdu.StepCommand(value),
			Qoc:   cmd.qoc,
		}
		if cmd.typeId == asdu.C_RC_TA_1 {
			asduCmd.Time = cmd.t
		}
		err = asdu.StepCmd(c.client104, cmd.typeId, coa, cmd.ca, asduCmd)
	case asdu.C_SE_NA_1, asdu.C_SE_TA_1:
		var value int16
		value, err = cast.ToInt16E(cmd.value)
		if err != nil {
			return err
		}
		asduCmd := asdu.SetpointCommandNormalInfo{
			Ioa:   cmd.ioa,
			Value: asdu.Normalize(value),
			Qos:   cmd.qos,
		}
		if cmd.typeId == asdu.C_SE_TA_1 {
			asduCmd.Time = cmd.t
		}
		err = asdu.SetpointCmdNormal(c.client104, cmd.typeId, coa, cmd.ca, asduCmd)
	case asdu.C_SE_NB_1, asdu.C_SE_TB_1:
		var value int16
		value, err = cast.ToInt16E(cmd.value)
		if err != nil {
			return err
		}
		asduCmd := asdu.SetpointCommandScaledInfo{
			Ioa:   cmd.ioa,
			Value: value,
			Qos:   cmd.qos,
		}
		if cmd.typeId == asdu.C_SE_TB_1 {
			asduCmd.Time = cmd.t
		}
		err = asdu.SetpointCmdScaled(c.client104, cmd.typeId, coa, cmd.ca, asduCmd)
	case asdu.C_SE_NC_1, asdu.C_SE_TC_1:
		var value float32
		value, err = cast.ToFloat32E(cmd.value)
		if err != nil {
			return err
		}
		asduCmd := asdu.SetpointCommandFloatInfo{
			Ioa:   cmd.ioa,
			Value: value,
			Qos:   cmd.qos,
		}
		if cmd.typeId == asdu.C_SE_TC_1 {
			asduCmd.Time = cmd.t
		}
		err = asdu.SetpointCmdFloat(c.client104, cmd.typeId, coa, cmd.ca, asduCmd)
	case asdu.C_BO_NA_1, asdu.C_BO_TA_1:
		var value uint32
		value, err = cast.ToUint32E(cmd.value)
		if err != nil {
			return err
		}
		asduCmd := asdu.BitsString32CommandInfo{
			Ioa:   cmd.ioa,
			Value: value,
		}
		if cmd.typeId == asdu.C_BO_TA_1 {
			asduCmd.Time = cmd.t
		}
		err = asdu.BitsString32Cmd(c.client104, cmd.typeId, coa, cmd.ca, asduCmd)
	default:
		err = fmt.Errorf("unknow type id %d", cmd.typeId)
	}

	return err
}

func activationCoa() asdu.CauseOfTransmission {
	return asdu.CauseOfTransmission{
		IsTest:     false,
		IsNegative: false,
		Cause:      asdu.Activation,
	}
}

// testConnect 测试端口连通性
func (c *Client) testConnect() error {
	url, _ := url.Parse(formatServerUrl(c.settings))
	var (
		conn net.Conn
		err  error
	)

	timeout := c.settings.Cfg104.ConnectTimeout0
	switch url.Scheme {
	case "tcps":
		conn, err = tls.DialWithDialer(&net.Dialer{Timeout: timeout}, "tcp", url.Host, c.settings.TLS)
	default:
		conn, err = net.DialTimeout("tcp", url.Host, timeout)
	}

	if err != nil {
		return err
	}
	return conn.Close()
}

func newClientOption(settings *Settings) *cs104.ClientOption {
	opts := cs104.NewOption()
	if settings.Cfg104 == nil {
		opts.SetConfig(cs104.DefaultConfig())
	}
	if settings.Params == nil {
		opts.SetParams(asdu.ParamsWide)
	}
	opts.SetAutoReconnect(settings.AutoConnect)
	opts.SetReconnectInterval(settings.ReconnectInterval)
	opts.SetTLSConfig(settings.TLS)

	server := formatServerUrl(settings)
	_ = opts.AddRemoteServer(server)
	return opts
}

func formatServerUrl(settings *Settings) string {
	var server string
	if settings.TLS != nil {
		server = "tcps://" + settings.Host + ":" + strconv.Itoa(settings.Port)
	} else {
		server = "tcp://" + settings.Host + ":" + strconv.Itoa(settings.Port)
	}
	return server
}
