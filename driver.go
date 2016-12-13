package app

import "github.com/murlokswarm/log"

var (
	driver Driver
)

// Driver is the interface that describes the implementation to handle platform
// specific rendering.
type Driver interface {
	Run()

	NewContext(ctx interface{}) Contexter

	AppMenu() Contexter

	Dock() Docker

	Resources() ResourcePath

	JavascriptBridge() string
}

// RegisterDriver registers the driver to be used when using the app package.
// Should be used only into a driver implementation, in an init function.
func RegisterDriver(d Driver) {
	driver = d
	log.Infof("driver %T is loaded", d)
}
