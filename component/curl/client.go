package curl

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/kenSevLeb/go-framework/component/trace"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Client struct {
	instance *http.Client

	options options
}

// 实例化
func New(opts ...Option) *Client {
	options := options{
		timeout:     time.Second * 10,
		traceHeader: "request-trace",
	}
	for _, opt := range opts {
		opt.apply(&options)
	}

	return &Client{
		instance: &http.Client{
			Transport: &http.Transport{ // 配置连接池
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				IdleConnTimeout: 30 * time.Second,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       options.timeout,
		},
		options: options,
	}
}

// New请求
func (cli *Client) NewRequest(method, url string, body io.Reader) *Request {
	req := &Request{
		Request: new(http.Request),
		err:     nil,
		client:  cli.instance,
	}
	request, err := http.NewRequest(method, url, body)
	if err == nil {
		request.Header.Set(cli.options.traceHeader, trace.Get())
		req.Request = request
	} else {
		req.err = err
	}
	return req
}

func (cli *Client) NewRequestWithHeader(method, url string, body io.Reader, header map[string]string) *Request {
	req := &Request{
		Request: new(http.Request),
		err:     nil,
		client:  cli.instance,
	}
	request, err := http.NewRequest(method, url, body)
	if err == nil {
		request.Header.Set(cli.options.traceHeader, trace.Get())
		for k, v := range header {
			request.Header.Set(k, v)
		}
		req.Request = request
	} else {
		req.err = err
	}
	return req
}

// 发送Get请求
func (cli *Client) Get(url string) (*response, error) {
	return cli.NewRequest(http.MethodGet, url, nil).Send()
}

// 发送Post请求
func (cli *Client) Post(url string, body io.Reader) (*response, error) {
	return cli.NewRequest(http.MethodPost, url, body).Send()
}

// 发送Put请求
func (cli *Client) Put(url string, body io.Reader) (*response, error) {
	return cli.NewRequest(http.MethodPut, url, body).Send()
}

// 发送Patch请求
func (cli *Client) Patch(url string, body io.Reader) (*response, error) {
	return cli.NewRequest(http.MethodPatch, url, body).Send()
}

// 发送删除请求
func (cli *Client) Delete(url string, body io.Reader) (*response, error) {
	return cli.NewRequest(http.MethodDelete, url, body).Send()
}

// 发送application/json格式的post请求
func (cli *Client) PostJson(url string, body io.Reader) (*response, error) {
	req := cli.NewRequest(http.MethodPost, url, body)
	if req.err != nil {
		return nil, req.err
	}
	req.Header.Set("Content-type", "application/json")
	return req.Send()
}

// 发送application/x-www-form-urlencoded格式的post请求
func (cli *Client) PostFormUrlEncode(url string, body io.Reader) (*response, error) {
	req := cli.NewRequest(http.MethodPost, url, body)
	if req.err != nil {
		return nil, req.err
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	return req.Send()
}

// 发送multipart/form-data格式的post请求
func (cli *Client) PostFormData(url string, body io.Reader) (*response, error) {
	req := cli.NewRequest(http.MethodPost, url, body)
	if req.err != nil {
		return nil, req.err
	}
	req.Header.Set("Content-type", "multipart/form-data")
	return req.Send()
}

// 上传文件
func (cli *Client) PostFile(url string, files map[string]string, params map[string]string) (*response, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for name, value := range params {
		_ = writer.WriteField(name, value)
	}

	for field, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("Open File: %v", err)
		}
		_ = file.Close()
		part, err := writer.CreateFormFile(field, filepath.Base(path))
		if err != nil {
			return nil, fmt.Errorf("Create Form File: %v", err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}
	}

	// 必须close
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("Close Writer: %v", err)
	}

	req := cli.NewRequest(http.MethodPost, url, body)
	if req.err != nil {
		return nil, req.err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req.Send()
}
