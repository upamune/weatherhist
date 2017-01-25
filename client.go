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

func newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	return nil, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
