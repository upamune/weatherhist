package weatherhist

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetDailyData(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	client, _ := NewClient(nil, nil)
	cases := []struct {
		Station      Station
		TargetDate   time.Time
		JSONFileName string
	}{
		{
			allStations["47570_36"],
			time.Date(2016, time.Month(6), 1, 0, 0, 0, 0, loc),
			"test/wakamatsu_daily_2016_06.json",
		},
		{
			allStations["0366_44"],
			time.Date(2015, time.Month(6), 1, 0, 0, 0, 0, loc),
			"test/hachioji_daily_2015_06.json",
		},
	}

	for _, c := range cases {
		daily, err := client.GetDailyData(c.Station, c.TargetDate)
		if err != nil {
			t.Errorf("can't get daily data: %s", err.Error())
			continue
		}

		dailyFromJSON := Daily{}
		b, err := ioutil.ReadFile(c.JSONFileName)
		if err != nil {
			t.Errorf("can't read daily from %s", c.JSONFileName)
			continue
		}
		err = json.Unmarshal(b, &dailyFromJSON)
		if err != nil {
			t.Errorf("can't parse %s: %s", c.JSONFileName, err.Error())
			continue
		}

		if len(daily.Days) != len(dailyFromJSON.Days) {
			t.Errorf("not equal length of daily\nExpected: %d, Actual: %d", len(dailyFromJSON.Days), len(daily.Days))
			continue
		}

		for i := 0; i < len(daily.Days); i++ {
			expectedDay := dailyFromJSON.Days[i]
			actualDay := daily.Days[i]

			if !reflect.DeepEqual(expectedDay, actualDay) {
				t.Errorf("Expected: %v, Actual:%v", expectedDay, actualDay)
			}
		}

	}
}
