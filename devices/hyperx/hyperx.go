package hyperx

import (
	"fmt"
	"log"
	"time"

	"github.com/Simaky/hid-v2"
	"github.com/getlantern/systray"

	"github.com/Simaky/Bathyx/resources"
)

const (
	vendorID      uint16 = 2385
	productID     uint16 = 5866
	deviceUsePage uint16 = 65299

	validResponseLength = 62
)

// LoadCloudFlightS loads 'HyperX Cloud Flight S'
func LoadCloudFlightS(menuItem *systray.MenuItem) {
	deviceInfo := hid.Enumerate(vendorID, productID)
	if len(deviceInfo) == 0 {
		log.Printf("Cloud Flight S does not connected")
	}

	var (
		device *hid.Device
		err    error
	)

	for idx := range deviceInfo {
		if deviceInfo[idx].UsagePage == deviceUsePage {
			device, err = deviceInfo[idx].Open()
			if err != nil {
				log.Printf("can't open device, err: %s", err)
				return
			}
		}
	}

	if device == nil {
		log.Printf("Cloud Flight S does not connected")
		return
	}

	for {
		askBatteryStatus(device)
		readBatteryStatus(device, menuItem)

		select {
		case <-time.Tick(time.Minute * 1): // update once per minute
			continue
		}
	}
}

func askBatteryStatus(device *hid.Device) {
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
		log.Printf("err: %s", err)
		return
	}
}

func readBatteryStatus(device *hid.Device, menuItem *systray.MenuItem) {
	buf := make([]byte, 2048)

	// TODO add timeout
	read, err := device.Read(buf)
	if err != nil {
		log.Printf("can't read device, err: %s", err)
		return
	}

	if read != validResponseLength {
		log.Printf("got invalid response")
		return
	}

	batteryPercentage := buf[7]

	if buf[5] != 0 && buf[6] != 0 && buf[3] == 2 {
		setBatteryPercent(menuItem, batteryPercentage)
	}
}

func setBatteryPercent(menuItem *systray.MenuItem, percentage byte) {
	title := fmt.Sprintf("HyperX Cloud Flight S: %d%% 🔋", percentage)
	systray.SetTooltip(title)
	menuItem.SetTitle(title)

	switch {
	case percentage <= 10:
		systray.SetIcon(resources.BatteryIcon10)
	case percentage <= 20:
		systray.SetIcon(resources.BatteryIcon20)
	case percentage <= 30:
		systray.SetIcon(resources.BatteryIcon30)
	case percentage <= 40:
		systray.SetIcon(resources.BatteryIcon40)
	case percentage <= 50:
		systray.SetIcon(resources.BatteryIcon50)
	case percentage <= 60:
		systray.SetIcon(resources.BatteryIcon60)
	case percentage <= 70:
		systray.SetIcon(resources.BatteryIcon70)
	case percentage <= 80:
		systray.SetIcon(resources.BatteryIcon80)
	case percentage <= 90:
		systray.SetIcon(resources.BatteryIcon90)
	case percentage <= 100:
		systray.SetIcon(resources.BatteryIcon100)
	}
}
