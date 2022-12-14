package bootstrap

import (
	"io"
)

type Provider struct{}

func (service Provider) Get(name string) string {
	switch name {

	// App Actions and Behaviors
	case "add":
		return service.get("plus")
	case "cancel":
		return service.get("x-lg")
	case "check-circle":
		return service.get("check-circle-fill")
	case "chevron-left":
		return service.get("chevron-left")
	case "chevron-right":
		return service.get("chevron-right")
	case "delete":
		return service.get("trash")
	case "edit":
		return service.get("pencil-square")
	case "email":
		return service.get("envelope")
	case "file":
		return service.get("file-earmark")
	case "flag":
		return service.get("flag-fill")
	case "folder":
		return service.get("folder")
	case "grip-vertical":
		return service.get("grip-vertical")
	case "grip-horizontal":
		return service.get("grip-horizontal")
	case "home":
		return service.get("house")
	case "info":
		return service.get("info-circle")
	case "link":
		return service.get("link-45deg")
	case "loading":
		return service.get("arrow-clockwise")
	case "login":
		return service.get("box-arrow-in-right")
	case "save":
		return service.get("check-lg")
	case "settings":
		return service.get("gear-fill")
	case "server":
		return service.get("hdd-stack")
	case "reply":
		return service.get("reply-fill")
	case "unlink":
		return service.get("link-45deg")
	case "upload":
		return service.get("cloud-arrow-up")
	case "user":
		return service.get("person-circle")
	case "users":
		return service.get("people")

	// Content Types
	case "article":
		return service.get("card-text")
	case "forward":
		return service.get("forward")
	case "inbox":
		return service.get("inbox")
	case "message":
		return service.get("chat-left-text")
	case "picture":
		return service.get("image")
	case "pictures":
		return service.get("images")
	case "shopping-cart":
		return service.get("cart3")

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
