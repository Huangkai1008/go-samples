package gee

import "log"

type RouterGroup struct {
	prefix      string        // add prefix to all routes in the group
	middlewares []HandlerFunc // support middlewares
	parent      *RouterGroup  // support group nesting
	engine      *Engine       // all groups share an Engine instance
}

// Group return new group with exist group
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// Get defines the method to add GET request
func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// Post defines the method to add POST request
func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
