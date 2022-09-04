package router

const (
	PublicPath = "/public"
	AdminPath  = "/admin"
	PingPath   = "/v1/ping"

	// BEGIN __INCLUDE_EXAMPLE__
	BaseSystemConfigPath = "/v1/config"
	// END __INCLUDE_EXAMPLE__
)

func PublicRoute(path string) string {
	return PublicPath + path
}

func AdminRoute(path string) string {
	return AdminPath + path
}
