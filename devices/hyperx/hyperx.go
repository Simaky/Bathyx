package hyperx

import (
	"context"
	"time"

	"github.com/Simaky/Bathyx/devices/types"
)

type Devices struct{}

// New returns struct with functions of all supported HyperX devices.
func New() *Devices {
	return &Devices{}
}

// CloudFlightS returns channel that sends actual battery percent
// and device status in provided timeout for 'HyperX Cloud Flight S' headphones.
// To stop fetching device information please use context.WithCancel().
func (d *Devices) CloudFlightS(ctx context.Context, timeout time.Duration) chan types.DeviceInfo {
	deviceInfoC := make(chan types.DeviceInfo)

	device, err := findDevice()
	if err != nil {
		go d.processError(err, deviceInfoC)
		return deviceInfoC
	}

	go func() {
		defer close(deviceInfoC)

		for {
			deviceInfoC <- getCloudFlightSInfo(device) // to run first time immediately

			select {
			case <-time.Tick(timeout):
				continue
			case <-ctx.Done():
				return
			}
		}
	}()

	return deviceInfoC
}

func (d *Devices) processError(err error, deviceInfoC chan<- types.DeviceInfo) {
	var deviceInfo types.DeviceInfo

	switch err {
	case types.ErrDeviceNotEnabled, types.ErrDeviceNotConnected:
		deviceInfo.Connected = false
	default:
		deviceInfo.Error = err
	}

	deviceInfoC <- deviceInfo
}
