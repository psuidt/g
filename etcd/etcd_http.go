package etcd

import (
	"time"

	"github.com/imdevlab/g"
	"github.com/valyala/fasthttp"
)

var HTTP = &Http{}

type Http struct{}

func (h *Http) Reqeust(app string, path string, method string, args *fasthttp.Args) (*fasthttp.Response, error) {
	req := &fasthttp.Request{}
	req.Header.SetMethod(method)

	s := GetServer(app)
	url := "http://" + s.IP + path
	switch method {
	case "GET":
		url = url + "?" + args.String()
	default:
		args.WriteTo(req.BodyWriter())
	}
	req.SetRequestURI(url)

	resp := &fasthttp.Response{}
	err := g.Cli.DoTimeout(req, resp, g.HTTP_REQ_TIMEOUT*time.Second)
	return resp, err
}
