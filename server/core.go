package server

import (
	"strconv"
	"sync"

	"github.com/wendy512/go-iecp5/asdu"
	"github.com/wendy512/go-iecp5/clog"
	"github.com/wendy512/go-iecp5/cs104"
)

// Settings 连接配置
type Settings struct {
	Host   string
	Port   int
	Cfg104 *cs104.Config //104协议规范配置
	Params *asdu.Params  //ASDU相关特定参数
	LogCfg *LogCfg
}

type LogCfg struct {
	Enable      bool //是否开启log
	LogProvider clog.LogProvider
}

type Server struct {
	settings              *Settings
	cs104Server           *cs104.Server
	connections           sync.Map // map[string]asdu.Connect
	connectionHandler     func(asdu.Connect)
	connectionLostHandler func(asdu.Connect)
}

func NewSettings() *Settings {
	cfg104 := cs104.DefaultConfig()
	return &Settings{
		Host:   "localhost",
		Port:   2404,
		Cfg104: &cfg104,
		Params: asdu.ParamsWide,
	}
}

func New(settings *Settings, handler CommandHandler) *Server {
	cs104Server := cs104.NewServer(&serverHandler{h: handler})
	cs104Server.SetConfig(*settings.Cfg104)
	cs104Server.SetParams(settings.Params)

	logCfg := settings.LogCfg
	if logCfg != nil {
		cs104Server.LogMode(logCfg.Enable)
		cs104Server.SetLogProvider(logCfg.LogProvider)
	}

	s := &Server{
		settings:    settings,
		cs104Server: cs104Server,
	}
	cs104Server.SetOnConnectionHandler(s.internalConnectionHandler)
	cs104Server.SetConnectionLostHandler(s.internalConnectionLostHandler)
	return s
}

func (s *Server) Start() {
	addr := s.settings.Host + ":" + strconv.Itoa(s.settings.Port)
	go s.cs104Server.ListenAndServer(addr)
}

func (s *Server) Stop() {
	_ = s.cs104Server.Close()
}

// SetOnConnectionHandler set on connect handler
func (s *Server) SetOnConnectionHandler(f func(asdu.Connect)) {
	s.connectionHandler = f
}

// SetConnectionLostHandler set connect lost handler
func (s *Server) SetConnectionLostHandler(f func(asdu.Connect)) {
	s.connectionLostHandler = f
}

// GetConnections get current connections
func (s *Server) GetConnections() []asdu.Connect {
	connects := make([]asdu.Connect, 0)
	s.connections.Range(func(key, value any) bool {
		connects = append(connects, value.(asdu.Connect))
		return true
	})
	return connects
}

func (s *Server) internalConnectionHandler(conn asdu.Connect) {
	addr := conn.UnderlyingConn().RemoteAddr().String()
	s.connections.Store(addr, conn)

	if s.connectionHandler != nil {
		s.connectionHandler(conn)
	}
}

func (s *Server) internalConnectionLostHandler(conn asdu.Connect) {
	addr := conn.UnderlyingConn().RemoteAddr().String()
	s.connections.Delete(addr)

	if s.connectionLostHandler != nil {
		s.connectionLostHandler(conn)
	}
}
