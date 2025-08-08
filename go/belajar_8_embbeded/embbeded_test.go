package belajar8embbeded

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string
func TestString(t *testing.T) {
	fmt.Println(version)
}