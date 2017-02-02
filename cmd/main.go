package main

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/upamune/weatherhist"
)

func main() {
	client, _ := weatherhist.NewClient(nil, nil)

	// 2010年1月の若松の気象データを取得する
	station, err := weatherhist.GetStation("47570", "36")
	if err != nil {
		panic(err)
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	t := time.Date(2010, time.Month(1), 1, 0, 0, 0, 0, loc)
	daily, err := client.GetDailyData(station, t)
	if err != nil {
		panic(err)
	}
	pp.Println(daily.Days[5])

}
