//go:build !windows && !darwin && !linux
// +build !windows,!darwin,!linux

package systheme

func GetTheme() (Theme, error) {
	return Unknown, ErrUnsupported
}

func GetSystemTheme() (Theme, error) {
	return Unknown, ErrUnsupported
}

func GetAppTheme() (Theme, error) {
	return Unknown, ErrUnsupported
}
