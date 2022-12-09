package logger

import (
	"fmt"
	"github.com/ryboe/q"
)

func Println(a ...any) {
	fmt.Println("[extended logger]: ", a)
	q.Q(a...)
}
