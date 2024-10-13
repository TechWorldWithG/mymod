package mymod

import "github.com/TechWorldWithG/mymod/internal/demo"

func External() string {
	return demo.Demo()
}
