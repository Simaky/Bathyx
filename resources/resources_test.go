package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBatteryIcon(t *testing.T) {
	tests := []struct {
		name          string
		percentage    int
		expectedError bool
	}{
		{
			name:          "100",
			percentage:    100,
			expectedError: false,
		},
		{
			name:          "Minus value",
			percentage:    -1234,
			expectedError: false,
		},
		{
			name:          "Not round",
			percentage:    57,
			expectedError: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			icon, err := GetBatteryIcon(testCase.percentage)
			if testCase.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NotNil(t, icon)
		})
	}
}
