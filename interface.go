package icon

import "io"

type Provider interface {
	Get(name string) string
	Write(name string, writer io.Writer)
}
