package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/alice"
)

func main() {
	err := InitializeSettings()
	if err != nil {
		fmt.Println("SYSTEM ERROR [FLIPPERPROX]: Error while initializing app settings -> ", err.Error())
		return
	}
	err = InitializeProxies()
	if err != nil {
		fmt.Println("SYSTEM ERROR [FLIPPERPROX]: Error while initializing reverse proxies -> ", err.Error())
		return
	}

	proxyPort := ":" + Settings.FlipperProxServiceListener
	if proxyPort == ":" {
		proxyPort = ":80"
	}

	handlerChain := alice.New(RecoverFromPanics, Certbot, Healthz, RedirectCheck, ProxyRouter).ThenFunc(NotFound)
	err = http.ListenAndServe(proxyPort, handlerChain)
	if err != nil {
		fmt.Println("SYSTEM ERROR [FLIPPERPROX]: Error while starting HTTP server -> ", err.Error())
		return
	}
}
