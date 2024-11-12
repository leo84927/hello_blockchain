package common

import (
	"crypto/tls"
	"hello_blockchain/config"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type _http struct {
	fc *fasthttp.Client
}

func NewHttp() *_http {
	return &_http{}
}

func (c *_http) Init() {
	c.fc = &fasthttp.Client{
		MaxConnsPerHost: 60000,
		TLSConfig:       &tls.Config{InsecureSkipVerify: true},
		ReadTimeout:     config.HttpTimeout,
		WriteTimeout:    config.HttpTimeout,
	}

	if config.HttpProxy != "" && config.HttpProxy != "0.0.0.0" {
		c.fc.Dial = fasthttpproxy.FasthttpHTTPDialer(config.HttpProxy)
	}
}

func (c *_http) Close() {
	if c.fc != nil {
		c.fc.CloseIdleConnections()
	}
}

func (c *_http) HttpRequest(url, method string, requestBody []byte, headers map[string]string) ([]byte, int, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.SetRequestURI(url)
	req.Header.SetMethod(method)

	switch method {
	case "POST", "PUT":
		req.SetBody(requestBody)
	case "GET":
		if requestBody != nil {
			req.SetBody(requestBody)
		}
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	err := c.fc.DoTimeout(req, resp, config.HttpTimeout)
	return resp.Body(), resp.StatusCode(), err
}
