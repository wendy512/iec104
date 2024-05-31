package client

import "github.com/thinkgos/go-iecp5/asdu"

// ClientASDUCall  is the interface of client handler
type ClientASDUCall interface {
	OnInterrogation(*asdu.ASDU) error
	OnCounterInterrogation(*asdu.ASDU) error
	OnRead(*asdu.ASDU) error
	OnTestCommand(*asdu.ASDU) error
	OnClockSync(*asdu.ASDU) error
	OnResetProcess(*asdu.ASDU) error
	OnDelayAcquisition(*asdu.ASDU) error
	OnASDU(*asdu.ASDU) error
}
