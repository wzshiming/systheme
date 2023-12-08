package main

import (
	"fmt"
	"os"

	"github.com/wzshiming/systheme"
)

func main() {
	t, err := systheme.GetTheme()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch t {
	case systheme.Light:
		fmt.Println("light")
	case systheme.Dark:
		fmt.Println("dark")
	default:
		fmt.Println("unknown")
	}
	return
}
