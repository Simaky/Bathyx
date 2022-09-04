package devices

import "github.com/Simaky/Bathyx/devices/hyperx"

type (
	Devices struct {
		HyperX *hyperx.Devices
	}
)

// New returns available devices.
func New() *Devices {
	return &Devices{
		HyperX: hyperx.New(),
	}
}
