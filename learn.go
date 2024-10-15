package mymod

import "github.com/TechWorldWithG/mymod/v2/internal/demo"

func External() string {
	return demo.Demo()
}
