package http_client

import (
    "net"
    "net/http"
    "time"
)

var (
    globalTransport *http.Transport

    MaxIdleConns        = 150
    MaxIdleConnsPerHost = 5

    defaultSetting = HttpSettings{
        timeout:   time.Second,
        transport: nil,
    }
)

func init() {
    globalTransport = GetGlobalTransport()
}

type HttpSettings = HttpClient

func (h *HttpSettings) SetDefaultTimeout(timeout time.Duration) *HttpSettings {
    if h.timeout != timeout {
        h.timeout = timeout
    }

    return h
}

func GetGlobalTransport() *http.Transport {
    transport := &http.Transport{
        Proxy: http.ProxyFromEnvironment,
        DialContext: (&net.Dialer{
            Timeout:   30 * time.Second,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        MaxIdleConns:          MaxIdleConns,
        MaxIdleConnsPerHost:   MaxIdleConnsPerHost,
        IdleConnTimeout:       30 * time.Second,
        TLSHandshakeTimeout:   10 * time.Second,
        ExpectContinueTimeout: 1 * time.Second,
    }

    return transport
}
