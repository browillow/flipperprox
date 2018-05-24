package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ProxyTarget is an enumeration of all standard reverse proxy targets
type ProxyTarget int

const (
	// FlipperApp is the proxy target that handles requests for the web app
	FlipperApp ProxyTarget = 0
	// FlipperAPI is the proxy target that handles requests for the web api
	FlipperAPI ProxyTarget = 1
)

func (target ProxyTarget) String() string {
	names := [...]string{"FlipperApp", "FlipperApi"}
	if target < FlipperApp || target > FlipperAPI {
		return "Unknown"
	}
	return names[target]
}

// ProxyHandlers contains pre-initialized handlers for reverse proxying to standard targets
var ProxyHandlers map[ProxyTarget]http.Handler

func buildReverseProxyHandler(host, port string) (http.Handler, error) {
	address := "http://" + host + ":" + port
	dest, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse proxy address: %v", address)
	}
	return httputil.NewSingleHostReverseProxy(dest), nil
}

// InitializeProxies builds the global map of pre-initialized reverse proxy handlers
func InitializeProxies() error {
	proxyDefinitions := make(map[ProxyTarget][]string)
	var buildProxyDefinition = func(host, port string) []string {
		return []string{host, port}
	}
	proxyDefinitions[FlipperApp] = buildProxyDefinition(Settings.FlipperAppServiceHost, Settings.FlipperAppServicePort)
	proxyDefinitions[FlipperAPI] = buildProxyDefinition(Settings.FlipperAPIServiceHost, Settings.FlipperAPIServicePort)

	ProxyHandlers = make(map[ProxyTarget]http.Handler)
	for key, value := range proxyDefinitions {
		handler, err := buildReverseProxyHandler(value[0], value[1])
		if err != nil {
			fmt.Println("ERROR [FLIPPERPROX]: Error while building reverse proxy: ", err.Error())
			ProxyHandlers[key] = http.HandlerFunc(NotFound)
		} else {
			ProxyHandlers[key] = handler
		}
	}
	return nil
}

// NotFound will return a 404 Not Found response
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
