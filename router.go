package cycapi

import (
	"net/http"
	"path"
	"strings"
)

type HandlerFunc func(*Context)

type version struct {
	Name		string
	Resources	map[string]resource
}

type resource struct {
	Name         string
	Handler      HandlerFunc
	Subresources map[string]subresource
}

type subresource struct {
	Name    string
	Handler HandlerFunc
}

type Routes struct {
	Map map[string]version
}

var Router *Routes

func init() {
	Router = &Routes{
		make(map[string]version),
	}

}

// Main entry point for each call
func (r *Routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	ctx := NewContext(w, req)
	defer ctx.Request.Body.Close()
	ctx.RouteInfo.Method = req.Method

	// Remove trailing slashes
	urlPath := path.Clean("/" + req.URL.Path)

	for place, value := range strings.Split(urlPath, "/") {
		switch place {
		case 1:
			ctx.RouteInfo.VersionName = value
			break
		case 2:
			// Always check for a custom method i.e. user:login
			resourceParts := strings.SplitN(value, ":", 2)

			// The resource will always be in position 0
			ctx.RouteInfo.ResourceName = resourceParts[0]

			// Add the custom method if it exists
			if len(resourceParts) == 2 {
				ctx.RouteInfo.CustomMethod = resourceParts[1]
			}

			break
		case 3:
			ctx.RouteInfo.ResourceID = value
			break
		case 4:
			// Always check for a custom method i.e. user:login
			resourceParts := strings.SplitN(value, ":", 2)

			// The resource will always be in position 0
			ctx.RouteInfo.SubresourceName = resourceParts[0]

			// Add the custom method if it exists
			if len(resourceParts) == 2 {
				ctx.RouteInfo.CustomMethod = resourceParts[1]
			}
			break
		case 5:
			ctx.RouteInfo.SubresourceID = value
			break
		default:
			// Ignore
			break
		}
	}

	// Load the subresource
	if ctx.RouteInfo.SubresourceName != "" {
		resource, exists := Router.Map[ctx.RouteInfo.VersionName].Resources[ctx.RouteInfo.ResourceName]
		if !exists {
			ctx.NotFound()
		} else {
			subresource, exists := resource.Subresources[ctx.RouteInfo.SubresourceName]
			if !exists {
				ctx.NotFound()
			} else {
				if ctx.RouteInfo.Method == "GET" && ctx.RouteInfo.SubresourceID == "" {
					ctx.RouteInfo.Method = "LIST"
				}
				subresource.Handler.ServeHTTP(ctx)
			}
		}
	} else {
		resource, exists := Router.Map[ctx.RouteInfo.VersionName].Resources[ctx.RouteInfo.ResourceName]
		if !exists {
			ctx.NotFound()
		} else {
			if ctx.RouteInfo.Method == "GET" && ctx.RouteInfo.ResourceID == "" {
				ctx.RouteInfo.Method = "LIST"
			}
			resource.Handler.ServeHTTP(ctx)
		}
	}
}

func (f HandlerFunc) ServeHTTP(context *Context) {
	f(context)
}

func (r *Routes) AddVersion(name string) {
	var version version
	version.Name = name
	version.Resources = make(map[string]resource)

	Router.Map[name] = version
}

func (r *Routes) AddResource(name string, versionName string, handler HandlerFunc) {
	var resourceObj resource
	resourceObj.Name = name
	resourceObj.Handler = handler
	resourceObj.Subresources = make(map[string]subresource)

	_, exists := Router.Map[versionName]
	if !exists {
		var version version
		version.Name = name
		version.Resources = make(map[string]resource)

		Router.Map[versionName] = version
	}

	Router.Map[versionName].Resources[name] = resourceObj
}

func (r *Routes) AddSubresource(name string, version string, resourceName string, handler HandlerFunc) {
	var subresource subresource
	subresource.Name = name
	subresource.Handler = handler

	Router.Map[version].Resources[resourceName].Subresources[name] = subresource
}
