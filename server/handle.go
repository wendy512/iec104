package server

import (
	"github.com/wendy512/go-iecp5/asdu"
	"time"
)

type serverHandler struct {
	h CommandHandler
}

func (s *serverHandler) InterrogationHandler(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierOfInterrogation) error {
	return s.h.OnInterrogation(conn, pack, quality)
}

func (s *serverHandler) CounterInterrogationHandler(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierCountCall) error {
	return s.h.OnCounterInterrogation(conn, pack, quality)
}

func (s *serverHandler) ReadHandler(conn asdu.Connect, pack *asdu.ASDU, addr asdu.InfoObjAddr) error {
	return s.h.OnRead(conn, pack, addr)
}

func (s *serverHandler) ClockSyncHandler(conn asdu.Connect, pack *asdu.ASDU, time time.Time) error {
	return s.h.OnClockSync(conn, pack, time)
}

func (s *serverHandler) ResetProcessHandler(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierOfResetProcessCmd) error {
	return s.h.OnResetProcess(conn, pack, quality)
}

func (s *serverHandler) DelayAcquisitionHandler(conn asdu.Connect, pack *asdu.ASDU, msec uint16) error {
	return s.h.OnDelayAcquisition(conn, pack, msec)
}

func (s *serverHandler) ASDUHandler(conn asdu.Connect, pack *asdu.ASDU) error {
	return s.h.OnASDU(conn, pack)
}
