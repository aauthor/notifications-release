package network

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var _transports map[bool]http.RoundTripper

func init() {
	_transports = map[bool]http.RoundTripper{
		true:  _buildTransport(true),
		false: _buildTransport(false),
	}
}

func buildTransport(skipVerifySSL bool) http.RoundTripper {
	return _transports[skipVerifySSL]
}

func _buildTransport(skipVerifySSL bool) http.RoundTripper {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipVerifySSL,
		},
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
}
