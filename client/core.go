package client

import (
	"crypto/tls"
	"github.com/thinkgos/go-iecp5/asdu"
	"github.com/thinkgos/go-iecp5/clog"
	"github.com/thinkgos/go-iecp5/cs104"
	"net"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	client104 *cs104.Client
	settings  *Settings
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

func New(settings *Settings) *Client {
	opts := newClientOption(settings)
	handler := &clientHandler{}
	client104 := cs104.NewClient(handler, opts)
	logCfg := settings.LogCfg
	if logCfg != nil {
		client104.LogMode(logCfg.Enable)
		client104.SetLogProvider(logCfg.LogProvider)
	}

	return &Client{
		client104: client104,
	}
}

func (c *Client) Connect() error {
	if err := c.testConnect(); err != nil {
		return err
	}
	return c.client104.Start()
}

func (c *Client) Close() error {
	return c.client104.Close()
}

func (c *Client) SetLogCfg(cfg LogCfg) {
	c.client104.LogMode(cfg.Enable)
	c.client104.SetLogProvider(cfg.LogProvider)
}

func (c *Client) SetOnConnectHandler(f func(c *Client)) {
	c.client104.SetOnConnectHandler(func(_ *cs104.Client) {
		f(c)
	})
}

func (c *Client) SetConnectionLostHandler(f func(c *Client)) {
	c.client104.SetOnConnectHandler(func(_ *cs104.Client) {
		f(c)
	})
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
	conn.Close()
	return nil
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
	opts.SetReconnectInterval(settings.ReconnectInterval)
	opts.SetTLSConfig(settings.TLS)

	server := formatServerUrl(settings)
	opts.AddRemoteServer(server)
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
