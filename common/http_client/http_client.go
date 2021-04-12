package http_client

import (
    "context"
    "net/http"
    "strings"
    "time"
)

type HttpClient struct {
    timeout       time.Duration
    transport     http.RoundTripper
    debug         bool
    client        *http.Client
    checkRedirect func(req *http.Request, via []*http.Request) error
}

type HttpClientOption func(client *HttpClient)

func WithTimeoutOpt(timeout time.Duration) HttpClientOption {
    return func(client *HttpClient) {
        client.SetDefaultTimeout(timeout)
    }
}

func NewHttpClient() *HttpClient {
    client := &HttpClient{}
    *client = defaultSetting

    return client
}

func (h *HttpClient) Request(ctx context.Context, rawUrl, method string) *HttpRequest {
    httpReq := &HttpRequest{
        method: strings.ToUpper(method),
        url:    rawUrl,
        body:   nil,
        header: nil,
    }

    if ctx == nil {
        ctx = context.Background()
    }
    httpReq.ctx = ctx

    h.checkClient()
    httpReq.httpClient = h

    return httpReq
}

func (h *HttpClient) checkClient() {
    if h.client != nil {
        return
    }

    trans := h.transport
    if trans == nil {
        trans = globalTransport
    } else {
        if v, ok := trans.(*http.Transport); ok {
            trans = v
        }
    }

    h.client = &http.Client{
        Transport:     trans,
        CheckRedirect: h.checkRedirect,
    }
}
