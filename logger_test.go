package logger

import (
	"testing"
)

func TestPrintln(t *testing.T) {
	type a struct {
		b string
	}
	Println("something", &a{})	
}