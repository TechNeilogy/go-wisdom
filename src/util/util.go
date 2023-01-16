package util

import (
	"fmt"
	"strings"
)

func PrintHeader(msg string) {
	fmt.Println("")
	fmt.Printf("+%v+\n", strings.Repeat("-", len(msg)+2))
	fmt.Printf("| %v |\n", msg)
	fmt.Printf("+%v+\n", strings.Repeat("-", len(msg)+2))
	fmt.Println("")
}
