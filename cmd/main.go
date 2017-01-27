package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/upamune/weatherhist"
)

func main() {
	client, _ := weatherhist.NewClient(nil, nil)

	dailies := []weatherhist.Daily{}

	// 若松の気象データを2010年1月から30ヶ月後まで取得して tenki.json に保存する
	ob, err := weatherhist.GetStation("1044", "36")
	if err != nil {
		panic(err)
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	t := time.Date(2010, time.Month(1), 1, 0, 0, 0, 0, loc)
	for i := 0; i < 30; i++ {
		daily, err := client.GetDailyData(ob, t)
		if err != nil {
			panic(err)
		}
		dailies = append(dailies, daily)
		t = t.AddDate(0, 1, 0)
	}

	fileName := "tenki.json"
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(f)
	err = enc.Encode(dailies)
	if err != nil {
		panic(err)
	}
}
