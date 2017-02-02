# weatherhist

[![CircleCI](https://circleci.com/gh/upamune/weatherhist/tree/master.svg?style=svg)](https://circleci.com/gh/upamune/weatherhist/tree/master)

Get weather hitory data from JMA Database.

[観測地点一覧](https://github.com/upamune/weatherhist/wiki/%E8%A6%B3%E6%B8%AC%E5%9C%B0%E7%82%B9%E4%B8%80%E8%A6%A7)

```go
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

// weatherhist.Day{
//   Date:          2010-01-06 00:00:00 Asia/Tokyo,
//   Precipitation: weatherhist.Precipitation{
//     Total:            &2.000000,
//     MaxPrecipitation: weatherhist.MaxPrecipitation{
//       Hourly:          &1.000000,
//       EveryTenMinutes: &0.500000,
//     },
//   },
//   Temperature: weatherhist.Temperature{
//     Average: &0.200000,
//     Highest: &3.200000,
//     Lowest:  &-1.800000,
//   },
//   Aerovane: weatherhist.Aerovane{
//     AverageWindSpeed: &2.800000,
//     MaxWindSpeed:     weatherhist.MaxWindSpeed{
//       Speed:     &6.900000,
//       Direction: &"西",
//     },
//     MaxInstantaneousSpeed: weatherhist.MaxInstantaneousSpeed{
//       Speed:     &11.800000,
//       Direction: &"西",
//     },
//     MostFrequentWindDirection: (*string)(nil),
//   },
//   HoursOfSunlight: &5.000000,
//   Snow:            weatherhist.Snow{
//     SnowFall: weatherhist.SnowFall{
//       Total: &8.000000,
//     },
//     DeepestSnow: weatherhist.DeepestSnow{
//       Value: &22.000000,
//     },
//   },
//   AirPressure: weatherhist.AirPressure{
//     FieldPressure: weatherhist.FieldPressure{
//       Average: &979.799988,
//     },
//     SeaSurfacePressure: weatherhist.SeaSurfacePressure{
//       Average: &1006.599976,
//     },
//   },
//   Humidity: weatherhist.Humidity{
//     Average: &83.000000,
//     Lowest:  &57.000000,
//   },
// }
```
