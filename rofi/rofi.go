package rofi

import (
	"fmt"
	"os"
	"strconv"
)

const (
	optionStart = "\x00"
	optionSep   = "\x1f"
)

type Env struct {
	// The argument passed to this
	Arg string
	// The value of ROFI_RETV
	CurrentState int
	// the value of ROFI_INFO
	Info string
	// the value of ROFI_DATA
	Data string
}

func GetEnv() Env {
	env := Env{
		Info: os.Getenv("ROFI_INFO"),
		Data: os.Getenv("ROFI_DATA"),
	}
	args := os.Args[1:]
	if len(args) > 0 {
		env.Arg = args[0]
	}
	state, err := strconv.Atoi(os.Getenv("ROFI_RETV"))
	if err != nil {
		state = 0
	}
	env.CurrentState = state
	return env
}

func anyToA(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case int:
		return strconv.Itoa(v)
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%f", v)
	}
}
