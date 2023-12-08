package systheme

import (
	"testing"
)

func TestGetTheme(t *testing.T) {
	theme, err := GetTheme()
	if err != nil {
		t.Error(err)
	}

	t.Log(theme)
}

func TestGetSystemTheme(t *testing.T) {
	theme, err := GetSystemTheme()
	if err != nil {
		t.Error(err)
	}

	t.Log(theme)
}

func TestGetAppTheme(t *testing.T) {
	theme, err := GetAppTheme()
	if err != nil {
		t.Error(err)
	}

	t.Log(theme)
}
