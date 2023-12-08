package systheme

// Theme represents the theme of the system.
//
//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=Theme
type Theme int

const (
	Unknown Theme = iota
	Light
	Dark
)
