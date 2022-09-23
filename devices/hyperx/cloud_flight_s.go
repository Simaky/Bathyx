package hyperx

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Simaky/hid-v2"

	"github.com/Simaky/Bathyx/devices/types"
)

const (
	vendorID      uint16 = 2385
	productID     uint16 = 5866
	deviceUsePage uint16 = 65299

	validResponseLength = 62

	defaultReadTimeout = time.Second * 10
	maxBufferSize      = 2048
)

// getCloudFlightSInfo returns 'Cloud Flight S' information.
func getCloudFlightSInfo(device *hid.Device) types.DeviceInfo {
	err := askBatteryStatus(device)
	if err != nil {
		return types.DeviceInfo{Error: err}
	}

	percentage, err := readBatteryPercentage(device)
	if err != nil {
		return types.DeviceInfo{Error: err}
	}

	return types.DeviceInfo{Connected: true, BatteryPercentage: percentage}
}

// findDevice search device though all connected devices, and checks does it enabled and connected.
func findDevice() (*hid.Device, error) {
	deviceInfo := hid.Enumerate(vendorID, productID)
	if len(deviceInfo) == 0 {
		return nil, types.ErrDeviceNotConnected
	}

	var (
		device *hid.Device
		err    error
	)

	for idx := range deviceInfo {
		if deviceInfo[idx].UsagePage == deviceUsePage {
			device, err = deviceInfo[idx].Open()
			if err != nil {
				return nil, fmt.Errorf("can't open device, err: %w", err)
			}
		}
	}

	if device == nil || !isEnabled(device) {
		return nil, types.ErrDeviceNotEnabled
	}

	return device, nil
}

// isEnabled try to read information from device to make sure
// that device enabled and correctly connected.
func isEnabled(device *hid.Device) bool {
	err := askBatteryStatus(device)
	if err != nil {
		return false
	}

	buf, read, err := readWithTimeout(device, defaultReadTimeout)
	if read == 0 || err != nil || len(buf) == 0 {
		return false
	}
	return true
}

// askBatteryStatus sends request battery percentage message to the device.
// nolint: gomnd
func askBatteryStatus(device *hid.Device) error {
	buffer := []byte{
		0x06,
		0x00,
		0x02,
		0x00,
		0x9a,
		0x00,
		0x00,
		0x68,
		0x4a,
		0x8e,
		0x0a,
		0x00,
		0x00,
		0x00,
		0xbb,
		0x02,
	}
	_, err := device.Write(buffer)
	if err != nil {
		err = fmt.Errorf("can't ask device status, err: %w", err)
	}

	return err
}

// readBatteryPercentage reads bytes from device and find battery percentage.
func readBatteryPercentage(device *hid.Device) (int, error) {
	buf, read, err := readWithTimeout(device, defaultReadTimeout)
	if err != nil {
		return 0, err
	}

	if read != validResponseLength {
		return 0, fmt.Errorf("got invalid response from device")
	}

	if buf[5] != 0 && buf[6] != 0 && buf[3] == 2 {
		return int(buf[7]), nil
	}

	return 0, fmt.Errorf("can't get battery percentage")
}

// isEnabled try to read information from device to make sure
// that device enabled and correctly connected.
func readWithTimeout(device *hid.Device, timeout time.Duration) ([]byte, int, error) {
	if device == nil {
		return nil, 0, fmt.Errorf("can't read from empty device")
	}

	var (
		read int
		err  error
	)

	buf := make([]byte, maxBufferSize)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()

		// TODO Update lib to support ctx with cancel
		read, err = device.Read(buf)
		if err != nil {
			err = fmt.Errorf("can't read from device, err: %w", err)
			log.Println(err)
			return
		}
	}()

	select {
	case <-ctx.Done():
		return buf, read, err
	case <-time.Tick(timeout):
		return nil, 0, types.ErrDeviceNotEnabled
	}
}
