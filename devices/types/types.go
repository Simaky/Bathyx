package types

import "fmt"

type (
	DeviceInfo struct {
		BatteryPercentage int
		Connected         bool
		Error             error
	}
)

var (
	ErrDeviceNotConnected = fmt.Errorf("device not connected")
	ErrDeviceNotEnabled   = fmt.Errorf("device not enabled")
)
