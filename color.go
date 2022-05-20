package fasthttpunit

import (
	"fmt"
)

// Red 错误
func Red(format string, args ...interface{}) string {
	return fmt.Sprintf("\033[0;31m"+format+"\033[0m", args...)
}

// Green 成功
func Green(format string, args ...interface{}) string {
	return fmt.Sprintf("\033[0;32m"+format+"\033[0m", args...)
}

// Yellow 警告
func Yellow(format string, args ...interface{}) string {
	return fmt.Sprintf("\033[0;33m"+format+"\033[0m", args...)
}
