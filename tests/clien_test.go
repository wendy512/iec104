package tests

import (
	"fmt"
	"github.com/thinkgos/go-iecp5/asdu"
	"github.com/wendy512/iec104/client"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	settings := client.NewSettings()
	settings.Host = "192.168.1.156"
	settings.LogCfg = &client.LogCfg{Enable: true}
	c := client.New(settings, &clientCall{})
	c.SetOnConnectHandler(func(c *client.Client) {
		// 连接成功以后做的操作
	})
	if err := c.Connect(); err != nil {
		t.Errorf("client connect error %v\n", err)
		t.FailNow()
	}

	time.Sleep(time.Second * 2)
	if err := c.SendInterrogationCmd(1); err != nil {
		t.Errorf("send interrogation cmd error %v\n", err)
		t.FailNow()
	}

	time.Sleep(time.Second * 2)
	if err := c.Close(); err != nil {
		t.Errorf("close error %v\n", err)
		t.FailNow()
	}
	time.Sleep(time.Second * 1)
}

type clientHandler struct {
}

type clientCall struct {
}

func (c *clientCall) OnInterrogation(asdu *asdu.ASDU) error {
	fmt.Println("OnInterrogation")
	return nil
}

func (c *clientCall) OnCounterInterrogation(asdu *asdu.ASDU) error {
	fmt.Println("OnCounterInterrogation")
	return nil
}

func (c *clientCall) OnRead(asdu *asdu.ASDU) error {
	fmt.Println("OnRead")
	return nil
}

func (c *clientCall) OnTestCommand(asdu *asdu.ASDU) error {
	fmt.Println("OnTestCommand")
	return nil
}

func (c *clientCall) OnClockSync(asdu *asdu.ASDU) error {
	fmt.Println("OnClockSync")
	return nil
}

func (c *clientCall) OnResetProcess(asdu *asdu.ASDU) error {
	fmt.Println("OnResetProcess")
	return nil
}

func (c *clientCall) OnDelayAcquisition(asdu *asdu.ASDU) error {
	fmt.Println("OnDelayAcquisition")
	return nil
}

func (c *clientCall) OnASDU(asdu *asdu.ASDU) error {
	fmt.Println("OnASDU")
	return nil
}
