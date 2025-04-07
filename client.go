package nga

import (
	"io"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html/charset"
)

type Client struct {
	client *resty.Client

	Config
}

func NewClient(config Config) *Client {
	return &Client{
		Config: config,
		client: resty.New(),
	}
}

type Config struct {
	BaseUrl        string
	NgaPassportUid string
	NgaPassportCid string
}

func (c *Client) getClient(opts ...func(*resty.Request)) *resty.Request {
	r := c.client.R()
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (c *Client) getDefaultClient() *resty.Request {
	return c.getClient(c.withCookie(), c.withOutput("11"))
}

func (c *Client) withCookie() func(*resty.Request) {
	return func(r *resty.Request) {
		r.
			SetHeader("User-Agent", "NGA_WP_JW").
			SetCookie(&http.Cookie{
				Name:  "nga_passport_uid",
				Value: c.NgaPassportUid,
			}).
			SetCookie(&http.Cookie{
				Name:  "nga_passport_cid",
				Value: c.NgaPassportCid,
			})

	}
}

// 为空时，默认为 11
// https://gitee.com/AgMonk/nga-api-doc#%E5%8F%91%E5%B8%96%E5%9B%9E%E5%A4%8D
func (c *Client) withOutput(method string) func(*resty.Request) {
	return func(r *resty.Request) {
		if method == "" {
			method = "11"
		}

		r.SetQueryParam("__output", "11")
	}
}

func gbkToUtf8(s string) (string, error) {
	gbkReader := strings.NewReader(s)

	utf8Reader, err := charset.NewReaderLabel("gbk", gbkReader)
	if err != nil {
		return "", err
	}

	utf8Data := new(strings.Builder)
	_, err = io.Copy(utf8Data, utf8Reader)
	if err != nil {
		return "", err
	}

	return utf8Data.String(), nil
}
