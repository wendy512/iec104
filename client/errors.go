package client

import "errors"

var (
	NotConnected = errors.New("the service request can not be executed because the client is not yet connected")
)
