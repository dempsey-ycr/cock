package cock

import (
	"html/template"
	"sync"
)

const defaultMultipartMemory = 32 << 10

// HandlerFunc defines the handler used by cock middleware as return value
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc array
type HandlersChain []HandlerFunc

// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// create an instance of Engine, by using New() or Default()
type Engine struct {
	RouterGroup

	// Enables automatic redirection if the current route can't be matched but a handler
	// for the path with(without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo,
	// the client is redirected to /foo with http status code 301 for GET
	// requests and 307 for all other request methods.
	RedirectTrailingSlash bool

	// If enabled, the router tries to fix the current requests path, if no
	// handle id registered for it.
	// First superfluous path elements like ../ or // are removed.
	// Afterwards the router does a case-insensitive lookup of the cleaned path.
	// If a handle can be found for this route, the router makes a redirection
	// to the corrected path with status code 301 for GET requests and 307 for
	// all other request methods.
	// For example /FOO and /..//Foo could be redirection to /foo.
	// RedirectTrailingSlash is independent of this option..
	RedirectFixedPath bool

	// If enabled, the router checks if another method is allowed for th current route,
	// if the current request can not be allowed.
	// If this is the case, the request is answered with 'Method Not Allowed' and HTTP status code 405
	// If no other Method is allowed, the request is delegated to the NotFound handler.
	HandleMethodNotAllowed bool
	ForwardedByClientIp    bool

	// #726 #755 if enabled, it will thrust some headers starting with
	// 'X-AppEngine...' for better integration with that PaaS(Platform-as-a-Server).
	AppEngine bool

	// If enabled, the url.RawPath will be used to find parameters.
	UseRawPath bool

	// If true, the path value will be unescaped.
	// If UseRawPath is false (by default), the UnescapePathValues effectively is true,
	// as url.Path gonna be used, which is already unescaped.
	UnescapePathValues bool

	// Value of 'maxMemory' param that is given to http.Request's ParseMultipartForm method call.
	MaxMultipartMemory int64

	delims render.Delime
	HTMLRender render.HTMLRender
	FuncMap template.FuncMap
	secureJsonPrefix string
	allNoRoute HandlersChain
	allNoMethod HandlersChain
	noRoute HandlersChain
	noMethod HandlersChain
	pool sync.Pool
	trees methodTrees
}
