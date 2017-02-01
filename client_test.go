package weatherhist

import (
	"testing"

	"log"

	"os"

	"time"

	"fmt"

	"net/url"

	"github.com/k0kubun/pp"
)

func TestNewClient(t *testing.T) {
	cases := []struct {
		URL         *string
		Logger      *log.Logger
		ExpectedURL string
	}{
		{
			nil,
			nil,
			defaultURL,
		},
		{
			getStringPointerAddress("https://google.com"),
			log.New(os.Stdout, "test", log.LstdFlags),
			"https://google.com",
		},
	}

	for _, c := range cases {
		client, err := NewClient(c.URL, c.Logger)
		if err != nil {
			t.Errorf("error occured when creating client with url: %v, logger: %v", c.URL, c.Logger)
		}
		if client.URL.String() != c.ExpectedURL {
			t.Errorf("Expected: %s, Actual: %s", c.ExpectedURL, client.URL.String())
		}
		if client.Logger == nil {
			t.Errorf("Expected: %v, Actual:%v", c.Logger, client.Logger)
		}
	}
}

func TestGetFullURL(t *testing.T) {
	client, _ := NewClient(nil, nil)
	loc, _ := time.LoadLocation("Asia/Tokyo")

	cases := []struct {
		SubPath     string
		Station     Station
		TargetDate  time.Time
		ExpectedURL string
	}{
		{
			"sample.php",
			Station{
				ID:          StationID("12345"),
				GroupNumber: "678",
				Name:        "test",
				Type:        StationTypeA,
			},
			time.Date(2016, time.Month(2), 1, 0, 0, 0, 0, loc),
			fmt.Sprintf("%s/%s",
				client.URL,
				"sample.php",
			),
		},
		{
			DailyPath,
			Station{
				ID:          StationID("12345"),
				GroupNumber: "678",
				Name:        "test",
				Type:        StationTypeA,
			},
			time.Date(2016, time.Month(2), 1, 0, 0, 0, 0, loc),
			fmt.Sprintf("%s/%s",
				client.URL,
				fmt.Sprintf(DailyPath, StationTypeA),
			),
		},
		{
			DailyPath,
			Station{
				ID:          StationID("12345"),
				GroupNumber: "678",
				Name:        "test",
				Type:        StationTypeS,
			},
			time.Date(2016, time.Month(2), 1, 0, 0, 0, 0, loc),
			fmt.Sprintf("%s/%s",
				client.URL,
				fmt.Sprintf(DailyPath, StationTypeS),
			),
		},
	}

	for _, c := range cases {
		urlStr := client.getFullURL(c.SubPath, c.Station, c.TargetDate)
		url, err := url.Parse(urlStr)
		if err != nil {
			t.Errorf("can't parse getFullURL: %s", err.Error())
		}

		urlWithOutQuery := url.Scheme + "://" + url.Host + url.Path
		if urlWithOutQuery != c.ExpectedURL {
			t.Errorf("Expected: %s, Actual: %s", c.ExpectedURL, urlWithOutQuery)
		}

		values := url.Query()
		if blockNo := values["block_no"][0]; blockNo != string(c.Station.ID) {
			t.Errorf("Expected: %s, Actual: %s", string(c.Station.ID), blockNo)
		}

		if precNo := values["prec_no"][0]; precNo != c.Station.GroupNumber {
			t.Errorf("Expected: %s, Actual: %s", c.Station.GroupNumber, precNo)
		}

		if view := values["view"][0]; view != "" {
			t.Errorf("Expected: %s, Actual: %s", "", view)
		}

		if year := values["year"][0]; year != fmt.Sprintf("%d", c.TargetDate.Year()) {
			t.Errorf("Expected: %d, Actual: %s", c.TargetDate.Year(), year)
		}

		if month := values["month"][0]; month != fmt.Sprintf("%d", int(c.TargetDate.Month())) {
			t.Errorf("Expected: %d, Actual: %s", int(c.TargetDate.Month()), month)
		}

		if day := values["day"][0]; day != fmt.Sprintf("%d", c.TargetDate.Day()) {
			t.Errorf("Expected: %d, Actual: %s", c.TargetDate.Day(), day)
		}

	}

}

func getFloat32PointerAddress(f32 float32) *float32 {
	return &f32
}
func TestGetFloatValue(t *testing.T) {
	cases := []struct {
		Case   string
		Expect *float32
	}{
		{
			"12.5",
			getFloat32PointerAddress(12.5),
		},
		{
			"10",
			getFloat32PointerAddress(10),
		},
		{
			"1.5]",
			getFloat32PointerAddress(1.5),
		},
		{
			"1.5)",
			getFloat32PointerAddress(1.5),
		},
		{
			"1.5 ",
			getFloat32PointerAddress(1.5),
		},
		{
			"///",
			nil,
		},
	}

	for _, c := range cases {
		f32 := getFloatValue(c.Case)
		if f32 == nil || c.Expect == nil {
			if f32 != c.Expect {
				t.Errorf("Expect: %v, Actual: %v", c.Expect, f32)
			}
			continue
		}
		if *f32 != *c.Expect {
			t.Errorf("Expect: %v, Actual: %v", *c.Expect, *f32)
		}
	}
}

func getStringPointerAddress(s string) *string {
	return &s
}

func TestGetStringValue(t *testing.T) {
	cases := []struct {
		Case   string
		Expect *string
	}{
		{
			"南南西",
			getStringPointerAddress("南南西"),
		},
		{
			"南南西 ",
			getStringPointerAddress("南南西"),
		},
		{
			"南南西)",
			getStringPointerAddress("南南西"),
		},
		{
			"南南西]",
			getStringPointerAddress("南南西"),
		},
		{
			NilValue,
			nil,
		},
		{
			"×",
			nil,
		},
		{
			"#",
			nil,
		},
	}

	for _, c := range cases {
		str := getStringValue(c.Case)
		if str == nil || c.Expect == nil {
			if str != c.Expect {
				pp.Println(str)
				t.Errorf("Expect: %v, Actual: %v", c.Expect, str)
			}
			continue
		}

		if *str != *c.Expect {
			t.Errorf("Expect: %v, Actual: %v", *c.Expect, *str)
		}
	}

}
