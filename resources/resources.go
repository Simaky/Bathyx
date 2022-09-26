package resources

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed images/logo.ico
	Logo []byte

	//go:embed images/headphones_icon.ico
	HeadphonesIcon []byte

	//go:embed images/battery/not_connected.ico
	HeadphonesNotConnectedIcon []byte

	//go:embed images/battery/10.ico
	batteryIcon10 []byte

	//go:embed images/battery/20.ico
	batteryIcon20 []byte

	//go:embed images/battery/30.ico
	batteryIcon30 []byte

	//go:embed images/battery/40.ico
	batteryIcon40 []byte

	//go:embed images/battery/50.ico
	batteryIcon50 []byte

	//go:embed images/battery/60.ico
	batteryIcon60 []byte

	//go:embed images/battery/70.ico
	batteryIcon70 []byte

	//go:embed images/battery/80.ico
	batteryIcon80 []byte

	//go:embed images/battery/90.ico
	batteryIcon90 []byte

	//go:embed images/battery/100.ico
	batteryIcon100 []byte

	batteries = map[int][]byte{
		10:  batteryIcon10,
		20:  batteryIcon20,
		30:  batteryIcon30,
		40:  batteryIcon40,
		50:  batteryIcon50,
		60:  batteryIcon60,
		70:  batteryIcon70,
		80:  batteryIcon80,
		90:  batteryIcon90,
		100: batteryIcon100,
	}
)

// GetBatteryIcon returns battery icon for given percentage (10-100).
func GetBatteryIcon(percentage int) ([]byte, error) {
	batteryIcon, ok := batteries[roundPercentage(percentage)]
	if !ok {
		return nil, fmt.Errorf("can't get battery icon, got percentage: %d", percentage)
	}
	return batteryIcon, nil
}

func roundPercentage(percentage int) int {
	if percentage <= 0 {
		return 10
	}
	if percentage > 100 {
		return 100
	}

	if percentage%10 == 0 {
		return percentage
	}

	return (10 - percentage%10) + percentage
}
