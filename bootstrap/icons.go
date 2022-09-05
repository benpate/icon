package bootstrap

import (
	"io"
)

type Provider struct{}

func (service Provider) Get(name string) string {
	switch name {
	case "add":
		return service.get("plus")
	case "cancel":
		return service.get("x-lg")
	case "delete":
		return service.get("trash")
	case "edit":
		return service.get("pencil-square")
	case "save":
		return service.get("check-lg")
	default:
		return service.get(name)
	}
}

func (service Provider) get(name string) string {
	return `<i class="bi bi-` + name + `"></i>`
}

func (service Provider) Write(name string, writer io.Writer) {
	writer.Write([]byte(service.Get(name)))
}
