package weatherhist

import (
	"log"
	"net/http"
	"net/url"

	"io/ioutil"
	"path"
	"strconv"
	"time"

	"fmt"
	"strings"

	"github.com/pkg/errors"
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

	c := http.DefaultClient
	return &Client{
		URL:        url,
		HTTPClient: c,
		Logger:     l,
	}, nil
}

func (c *Client) getFullURL(spath string, s Station, targetDate time.Time) string {
	q := url.Values{}
	q.Add("block_no", string(s.ID))
	q.Add("prec_no", s.GroupNumber)
	q.Add("year", strconv.Itoa(targetDate.Year()))
	q.Add("month", strconv.Itoa(int(targetDate.Month())))
	q.Add("day", strconv.Itoa(targetDate.Day()))
	q.Add("view", "")

	u := *c.URL
	u.RawQuery = q.Encode()

	switch spath {
	case DailyPath:
		spath = fmt.Sprintf(DailyPath, s.Type)
	}
	u.Path = path.Join(c.URL.Path, spath)
	return u.String()
}

func getFloatValueWithQuality(value string) FloatWithQuality {
	fwq := FloatWithQuality{}
	value = strings.TrimRight(value, " ")
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		fwq.IsBadQuality = true
		// "2.5 ]" とかをパースやってみる
		v := strings.TrimRight(value, ")]")
		f, err = strconv.ParseFloat(v, 32)
		if err != nil {
			return FloatWithQuality{
				Value: nil,
				IsBadQuality: true,
			}
		}
	}
	f32 := float32(f)
	fwq.Value = &f32

	return fwq
}

const NilValue = "///"

func getStringValueWithQuality(value string) StringWithQuality {
	swq := StringWithQuality{}
	value = strings.TrimRight(value, " ")

	if strings.ContainsAny(value, ")]") {
		swq.IsBadQuality = true
	}

	value = strings.TrimRight(value, ")]")
	swq.Value = &value

	if value == NilValue || value == "×" || value == "#" {
		swq.Value = nil
		swq.IsBadQuality = true
	}

	return swq
}
