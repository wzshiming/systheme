package systheme

import (
	"golang.org/x/sys/windows/registry"
)

func getTheme() (Theme, error) {
	return getSystemTheme()
}

func getSystemTheme() (Theme, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`, registry.QUERY_VALUE)
	if err != nil {
		return Unknown, err
	}
	defer k.Close()

	v, _, err := k.GetIntegerValue("SystemUsesLightTheme")
	if err != nil {
		return Unknown, err
	}

	if v == 0 {
		return Dark, nil
	}

	return Light, nil
}

func getAppTheme() (Theme, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`, registry.QUERY_VALUE)
	if err != nil {
		return Unknown, err
	}
	defer k.Close()

	v, _, err := k.GetIntegerValue("AppsUseLightTheme")
	if err != nil {
		return Unknown, err
	}

	if v == 0 {
		return Dark, nil
	}

	return Light, nil
}
