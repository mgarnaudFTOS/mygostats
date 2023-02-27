package mock_test

import (
	stats "github.com/mgarnaudFTOS/gostats"
	"github.com/mgarnaudFTOS/gostats/mock"
)

var (
	_ stats.Sink          = (*mock.Sink)(nil)
	_ stats.FlushableSink = (*mock.Sink)(nil)
)
