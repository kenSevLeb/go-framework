package curl

import (
	"fmt"
	"kenSevLeb/go-framework/util/bytes"
	"net/http"
)

// DefaultAdapter
var adapter = bytes.NewAdapter()

type Request struct {
	*http.Request

	err    error
	client *http.Client
}

// Send send http request
func (impl *Request) Send() (*response, error) {
	if impl.err != nil {
		return nil, impl.err
	}

	resp, err := impl.client.Do(impl.Request)
	if resp != nil { // 重定向错误时，resp不为nil
		// close response
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("no response")
	}

	res, err := adapter.Read(resp.Body)
	if err != nil {
		return nil, err
	}
	return &response{header: resp.Header, body: []byte(res), statusCode: resp.StatusCode}, nil
}

// 自定义Client
func (impl *Request) SetClient(client *http.Client) {
	impl.client = client
}

func (impl *Request) Err() error {
	return impl.err
}
