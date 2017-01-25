package main

import (
	"github.com/upamune/weatherhist"
	"time"
)

func main() {
	client, _ := weatherhist.NewClient(nil, nil)
	key := weatherhist.GetObservatoriesKey("1519", "69")
	client.GetDailyData(key, time.Now(), time.Now())

}
