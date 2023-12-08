package systheme

import (
	"fmt"
	"strings"
)

func getTheme() (Theme, error) {
	out, err := command("gsettings", "get", "org.gnome.desktop.interface", "gtk-theme")
	if err != nil {
		return Unknown, fmt.Errorf("%s: %w", out, err)
	}

	if strings.Contains(strings.ToLower(out), "dark") {
		return Dark, nil
	}

	return Light, nil
}

func getSystemTheme() (Theme, error) {
	return getTheme()
}

func getAppTheme() (Theme, error) {
	return getTheme()
}
