package router

type Router struct {
	Get    []Route
	Post   []Route
	Put    []Route
	Delete []Route
}

type Route struct {
	Method  string
	Path    string
	Handler string
}
