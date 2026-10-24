// Package icon gives a web application one interface for drawing icons, so that the icon
// library behind it can change without touching the code that draws them.
//
// Provider is that interface. Get returns the HTML markup for an icon by name, and Write
// writes the same markup to an io.Writer. Names are generic ("add", "delete", "user"), and
// each implementation translates them into its own library's icon names.
//
// The bootstrap subpackage implements Provider with Bootstrap Icons, and is currently the
// only implementation.
package icon
