package weatherhist

import (
	"time"

	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const DailyPath = "daily_a1.php"

type Daily struct {
	Days []Day
}

type Day struct {
	Date            time.Time
	Precipitation   Precipitation
	Temperature     Temperature
	Aerovane        Aerovane
	HoursOfSunlight string
	Snow            Snow
}

type Precipitation struct {
	Total            string
	MaxPrecipitation MaxPrecipitation
}

type MaxPrecipitation struct {
	Hourly          string
	EveryTenMinutes string
}

type Temperature struct {
	Average string
	Highest string
	Lowest  string
}

type Aerovane struct {
	AverageWindSpeed          string
	MaxWindSpeed              MaxWindSpeed
	MaxInstantaneousSpeed     MaxInstantaneousSpeed
	MostFrequentWindDirection string
}

type MaxWindSpeed struct {
	Speed     string
	Direction string
}

type MaxInstantaneousSpeed struct {
	Speed     string
	Direction string
}

type Snow struct {
	SnowFall    SnowFall
	DeepestSnow DeepestSnow
}

type SnowFall struct {
	Total string
}

type DeepestSnow struct {
	Value string
}

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = loc
}

func (c *Client) GetDailyData(id ObservationID, begin, end time.Time) (Daily, error) {
	daily := Daily{}

	c.getDailyDataFromPage()

	return daily, nil
}

func (c *Client) getDailyDataFromPage() (Daily, error) {
	url := c.getFullURL(DailyPath)
	year := 2016
	month := time.Month(1)
	daily := Daily{}
	doc, err := goquery.NewDocument(url + "?prec_no=69&block_no=1519&year=2016&month=1&day=1&view=")
	if err != nil {
		return daily, err
	}
	doc.Find("#tablefix1 > tbody > tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 || i == 1 || i == 2 {
			return
		}
		day := Day{}
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				date, _ := strconv.Atoi(s.Text())
				t := time.Date(year, month, date, 0, 0, 0, 0, time.Local)
				day.Date = t
			case 1:
				day.Precipitation.Total = s.Text()
			case 2:
				day.Precipitation.MaxPrecipitation.Hourly = s.Text()
			case 3:
				day.Precipitation.MaxPrecipitation.EveryTenMinutes = s.Text()
			case 4:
				day.Temperature.Average = s.Text()
			case 5:
				day.Temperature.Highest = s.Text()
			case 6:
				day.Temperature.Lowest = s.Text()
			case 7:
				day.Aerovane.AverageWindSpeed = s.Text()
			case 8:
				day.Aerovane.MaxWindSpeed.Speed = s.Text()
			case 9:
				day.Aerovane.MaxWindSpeed.Direction = s.Text()
			case 10:
				day.Aerovane.MaxInstantaneousSpeed.Speed = s.Text()
			case 11:
				day.Aerovane.MaxInstantaneousSpeed.Direction = s.Text()
			case 12:
				day.Aerovane.MostFrequentWindDirection = s.Text()
			case 13:
				day.HoursOfSunlight = s.Text()
			case 14:
				day.Snow.SnowFall.Total = s.Text()
			case 15:
				day.Snow.DeepestSnow.Value = s.Text()
				daily.Days = append(daily.Days, day)
			default:
				return
			}
		})
	})

	return daily, nil
}
