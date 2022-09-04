package main

import (
	"os"

	"github.com/getlantern/systray"

	"github.com/Simaky/Bathyx/devices/hyperx"
	"github.com/Simaky/Bathyx/resources"
)

const appName = "Bathyx™"

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(resources.HeadphonesNotConnectedIcon)
	systray.SetTitle("Bathyx")
	systray.SetTooltip("Waiting for headphones...")

	systray.AddMenuItem(appName, "https://github.com/Simaky/Bathyx").Disable()
	systray.AddSeparator()

	item := systray.AddMenuItem("HyperX Cloud Flight S ", "Quit the whole app")
	item.SetIcon(resources.HeadphonesIcon)
	item.Disable()

	systray.AddSeparator()
	exit := systray.AddMenuItem("Quit", "Quit the whole app")

	go hyperx.LoadCloudFlightS(item)

	go func() {
		select {
		case <-exit.ClickedCh:
			os.Exit(0)
		}
	}()
}

func onExit() {
	os.Exit(0)
}
