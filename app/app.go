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

// timeout for CPU usage optimisation.
const timeout = time.Second * 5

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
	systray.SetIcon(resources.Logo)
	// TODO show always if connected more than 1 device
	systray.SetTitle("Bathyx")
	systray.SetTooltip("Waiting for headphones...")

	appItem := systray.AddMenuItem(t.appName, "https://github.com/Simaky/Bathyx")
	appItem.SetIcon(resources.Logo)
	appItem.Disable()
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

	loadDevices(context.Background(), devices.New(), item)
}

func (*TrayApp) onExit() {}

// nolint:gosimple
func loadDevices(ctx context.Context, d *devices.Devices, item *systray.MenuItem) {
	cloudFlightSC := d.HyperX.CloudFlightS(ctx, timeout)

	// TODO remove nolint after adding more devices to select
	for {
		select {
		case resp := <-cloudFlightSC:
			processDeviceInfo(resp, item)
		}
	}
}

func processDeviceInfo(deviceInfo types.DeviceInfo, item *systray.MenuItem) {
	if deviceInfo.Error != nil {
		systray.SetIcon(resources.Logo)
		systray.SetTooltip("Waiting for headphones...")
		item.SetTitle("HyperX Cloud Flight S: waiting...")
		log.Println(deviceInfo.Error)
		return
	}

	if !deviceInfo.Connected {
		systray.SetIcon(resources.Logo)
		systray.SetTooltip("Waiting for headphones...")
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

	batteryIcon, err := resources.GetBatteryIcon(percentage)
	if err != nil {
		log.Println(err)
		return
	}
	systray.SetIcon(batteryIcon)
}
