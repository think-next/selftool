package http_client

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
)

var (
    HeaderContentType = "Content-Type"

    ContentTypeForm = "application/x-www-form-urlencoded"
    ContentTypeJson = "application/json; charset=utf-8"
)

type HttpRequest struct {
    method     string
    url        string
    body       io.Reader
    header     http.Header
    ctx        context.Context
    error      error
    httpClient *HttpClient
}

func (b *HttpRequest) WithBody(data interface{}) *HttpRequest {
    switch t := data.(type) {
    case string:
        b.body = bytes.NewReader([]byte(t))
    case []byte:
        b.body = bytes.NewReader(t)
    case url.Values:
        b.body = bytes.NewReader([]byte(t.Encode()))
        b.SetHeaderField(HeaderContentType, ContentTypeForm)
    case nil:
        b.body = nil
    default:
        withBody, err := json.Marshal(data)
        if err != nil {
            b.error = err
            return b
        }

        b.body = bytes.NewReader(withBody)
        b.SetHeaderField(HeaderContentType, ContentTypeJson)
    }

    return b
}

func (b *HttpRequest) SetHeaderField(key, value string) *HttpRequest {
    if b.header == nil {
        b.header = make(http.Header)
    }

    b.header.Set(key, value)
    return b
}

// need caller to close response
func (b *HttpRequest) Response(ctx context.Context) (*http.Response, error) {
    request, err := b.request(ctx)
    if err != nil {
        return nil, err
    }

    return b.do(ctx, request)
}

func (b *HttpRequest) BytesWithStatus(ctx context.Context) (int, []byte, error) {

    response, err := b.Response(ctx)
    if err != nil {
        return 0, nil, err
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return response.StatusCode, nil, err
    }

    return response.StatusCode, body, nil
}

func (b *HttpRequest) ToJSON(ctx context.Context, v interface{}) (int, error) {
    code, data, err := b.BytesWithStatus(ctx)
    if err != nil {
        return code, err
    }

    if err = json.Unmarshal(data, v); err == nil {
        return code, nil
    }

    return code, fmt.Errorf("to json error: %v, content: %s", err, string(data))
}

func (b *HttpRequest) request(ctx context.Context) (*http.Request, error) {
    req, err := http.NewRequestWithContext(ctx, b.method, b.url, b.body)
    if err != nil {
        return nil, err
    }

    // TODO inject trace
    req.Header = b.header
    return req, nil
}

func (b *HttpRequest) do(ctx context.Context, req *http.Request) (*http.Response, error) {
    return b.httpClient.client.Do(req)
}
