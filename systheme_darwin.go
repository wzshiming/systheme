package systheme

import (
	"fmt"
	"strings"
)

func getTheme() (Theme, error) {
	out, err := command("defaults", "read", "-g", "AppleInterfaceStyle")
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			return Light, nil
		}
		return Unknown, fmt.Errorf("%s: %w", out, err)
	}
	switch strings.TrimSpace(out) {
	case "Dark":
		return Dark, nil
	default:
		return Light, nil
	}
}

func getSystemTheme() (Theme, error) {
	return getTheme()
}

func getAppTheme() (Theme, error) {
	return getTheme()
}
