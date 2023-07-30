package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewOpenAIProxy() *httputil.ReverseProxy {
	remoteTarget, _ := url.Parse("https://api.openai.com")

	proxy := httputil.NewSingleHostReverseProxy(remoteTarget)
	proxy.Director = func(req *http.Request) {
		rewriteRequestURL(req, remoteTarget)
	}

	return proxy
}
