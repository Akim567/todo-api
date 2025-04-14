package log

import (
	"fmt"
	"os"
)

func Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("❌", msg)
	os.Exit(1)
}
