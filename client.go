package weatherhist

import (
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"io/ioutil"
	"context"
	"io"
	"encoding/json"
	"path"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *log.Logger
}

const defaultURL = "http://www.data.jma.go.jp/obd/stats/etrn/view"

func NewClient(urlStrp *string, logger *log.Logger) (*Client, error) {
	urlStr := defaultURL
	if urlStrp != nil {
		urlStr = *urlStrp
	}
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url")
	}

	l := log.New(ioutil.Discard, "", log.LstdFlags)
	if logger != nil {
		l = logger
	}

	c := &http.Client{}
	return &Client{
		URL:        url,
		HTTPClient: c,
		Logger:     l,
	}, nil
}

func (c *Client)newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.getFullURL(spath), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) getFullURL(spath string) string {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	return u.String()
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
