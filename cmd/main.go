package main

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/upamune/weatherhist"
)

func main() {
	client, _ := weatherhist.NewClient(nil, nil)
	observation, err := weatherhist.GetObservatory("1519", "69")
	if err != nil {
		panic(err)
	}
	daily, err := client.GetDailyData(observation, time.Now())
	if err != nil {
		panic(err)
	}
	pp.Println(daily)
}
