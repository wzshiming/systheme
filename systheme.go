package systheme

import (
	"fmt"
)

var ErrUnsupported = fmt.Errorf("unsupported platform")

// GetTheme returns the current theme of the system.
func GetTheme() (Theme, error) {
	return getTheme()
}

// GetSystemTheme returns the current theme of the system.
func GetSystemTheme() (Theme, error) {
	return getSystemTheme()
}

// GetAppTheme returns the current theme of the system.
func GetAppTheme() (Theme, error) {
	return getAppTheme()
}
