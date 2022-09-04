package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getlantern/systray"

	"github.com/Simaky/Bathyx/devices"
	"github.com/Simaky/Bathyx/devices/types"
	"github.com/Simaky/Bathyx/resources"
)

type TrayApp struct {
	appName string
}

// RunTray runs tray application.
func RunTray(appName string) {
	app := TrayApp{
		appName: appName,
	}
	systray.Run(app.onReady, app.onExit)
}

func (t *TrayApp) onReady() {
	systray.SetIcon(resources.HeadphonesNotConnectedIcon)
	// TODO show always if connected more than 1 device
	systray.SetTitle("Bathyx")
	systray.SetTooltip("Waiting for headphones...")

	systray.AddMenuItem(t.appName, "https://github.com/Simaky/Bathyx").Disable()
	systray.AddSeparator()

	// TODO show only if 1 device is connected
	item := systray.AddMenuItem("HyperX Cloud Flight S: waiting...", "")
	item.SetIcon(resources.HeadphonesIcon)
	item.Disable()

	systray.AddSeparator()
	exit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-exit.ClickedCh
		os.Exit(0)
	}()

	go loadDevices(context.Background(), devices.New(), item)
}

func (*TrayApp) onExit() {}

func loadDevices(ctx context.Context, d *devices.Devices, item *systray.MenuItem) {
	for {
		processDeviceInfo(<-d.HyperX.CloudFlightS(ctx, time.Minute*1), item)
	}
}

func processDeviceInfo(deviceInfo types.DeviceInfo, item *systray.MenuItem) {
	if deviceInfo.Error != nil {
		log.Println(deviceInfo.Error)
		return
	}

	if !deviceInfo.Connected {
		item.SetTitle("HyperX Cloud Flight S: waiting...")
		return
	}
	setBatteryPercent(item, deviceInfo.BatteryPercentage)
}

// nolint: gomnd
func setBatteryPercent(menuItem *systray.MenuItem, percentage int) {
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
