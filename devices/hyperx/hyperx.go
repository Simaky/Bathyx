package hyperx

import (
	"context"
	"errors"
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

	go func() {
		defer close(deviceInfoC)

		for {
			// to run first time immediately
			device, err := findDevice()
			if err != nil {
				go d.processError(err, deviceInfoC)
				continue
			}
			deviceInfoC <- getCloudFlightSInfo(device)
			device.Close()
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

func (*Devices) processError(err error, deviceInfoC chan<- types.DeviceInfo) {
	var deviceInfo types.DeviceInfo

	if errors.Is(err, types.ErrDeviceNotEnabled) || errors.Is(err, types.ErrDeviceNotConnected) {
		deviceInfo.Connected = false
	} else {
		deviceInfo.Error = err
	}

	deviceInfoC <- deviceInfo
}
