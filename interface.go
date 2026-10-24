package icon

import "io"

// Provider draws icons by name
type Provider interface {

	// Get returns the HTML markup for the named icon
	Get(name string) string

	// Write writes the HTML markup for the named icon to the writer
	Write(name string, writer io.Writer)
}
