package weatherhist

import (
	"time"

	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const DailyPath = "daily_%s1.php"

type Daily struct {
	Days []Day
}

type Day struct {
	Date            time.Time
	Precipitation   Precipitation
	Temperature     Temperature
	Aerovane        Aerovane
	HoursOfSunlight *float32
	Snow            Snow
	AirPressure     AirPressure
	Humidity        Humidity
}

type Precipitation struct {
	Total            *float32
	MaxPrecipitation MaxPrecipitation
}

type MaxPrecipitation struct {
	Hourly          *float32
	EveryTenMinutes *float32
}

type Temperature struct {
	Average *float32
	Highest *float32
	Lowest  *float32
}

type Aerovane struct {
	AverageWindSpeed          *float32
	MaxWindSpeed              MaxWindSpeed
	MaxInstantaneousSpeed     MaxInstantaneousSpeed
	MostFrequentWindDirection *string
}

type MaxWindSpeed struct {
	Speed     *float32
	Direction *string
}

type MaxInstantaneousSpeed struct {
	Speed     *float32
	Direction *string
}

type Snow struct {
	SnowFall    SnowFall
	DeepestSnow DeepestSnow
}

type SnowFall struct {
	Total *float32
}

type DeepestSnow struct {
	Value *float32
}

type AirPressure struct {
	FieldPressure      FieldPressure
	SeaSurfacePressure SeaSurfacePressure
}

type FieldPressure struct {
	Average *float32
}

type SeaSurfacePressure struct {
	Average *float32
}

type Humidity struct {
	Average *float32
	Lowest  *float32
}

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = loc
}

func (c *Client) GetDailyData(ob Station, targetDate time.Time) (Daily, error) {
	daily, err := c.getDailyDataFromPage(ob, targetDate)
	if err != nil {
		return daily, err
	}
	return daily, nil
}

func (c *Client) getDailyDataFromPage(st Station, targetDate time.Time) (Daily, error) {
	url := c.getFullURL(DailyPath, st, targetDate)
	daily := Daily{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return daily, err
	}
	switch st.Type {
	case StationTypeS:
		daily, err = getDailyDataFromPageTypeS(st, targetDate, doc)
	case StationTypeA:
		daily, err = getDailyDataFromPageTypeA(st, targetDate, doc)
	default:
		return daily, errors.Wrapf(err, "unkown station type: %s", st.Type)
	}

	return daily, nil
}

func getDailyDataFromPageTypeS(st Station, targetDate time.Time, doc *goquery.Document) (Daily, error) {
	daily := Daily{}
	doc.Find("#tablefix1 > tbody > tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 || i == 1 || i == 2 {
			return
		}
		day := Day{}
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				date, _ := strconv.Atoi(s.Text())
				t := time.Date(targetDate.Year(), targetDate.Month(), date, 0, 0, 0, 0, time.Local)
				day.Date = t
			case 1:
				day.AirPressure.FieldPressure.Average = getFloatValue(s.Text())
			case 2:
				day.AirPressure.SeaSurfacePressure.Average = getFloatValue(s.Text())
			case 3:
				day.Precipitation.Total = getFloatValue(s.Text())
			case 4:
				day.Precipitation.MaxPrecipitation.Hourly = getFloatValue(s.Text())
			case 5:
				day.Precipitation.MaxPrecipitation.EveryTenMinutes = getFloatValue(s.Text())
			case 6:
				day.Temperature.Average = getFloatValue(s.Text())
			case 7:
				day.Temperature.Highest = getFloatValue(s.Text())
			case 8:
				day.Temperature.Lowest = getFloatValue(s.Text())
			case 9:
				day.Humidity.Average = getFloatValue(s.Text())
			case 10:
				day.Humidity.Lowest = getFloatValue(s.Text())
			case 11:
				day.Aerovane.AverageWindSpeed = getFloatValue(s.Text())
			case 12:
				day.Aerovane.MaxWindSpeed.Speed = getFloatValue(s.Text())
			case 13:
				day.Aerovane.MaxWindSpeed.Direction = getStringValue(s.Text())
			case 14:
				day.Aerovane.MaxInstantaneousSpeed.Speed = getFloatValue(s.Text())
			case 15:
				day.Aerovane.MaxInstantaneousSpeed.Direction = getStringValue(s.Text())
			case 16:
				day.HoursOfSunlight = getFloatValue(s.Text())
			case 17:
				day.Snow.SnowFall.Total = getFloatValue(s.Text())
			case 18:
				day.Snow.DeepestSnow.Value = getFloatValue(s.Text())
				daily.Days = append(daily.Days, day)
			default:
				return
			}
		})
	})

	return daily, nil
}

func getDailyDataFromPageTypeA(st Station, targetDate time.Time, doc *goquery.Document) (Daily, error) {
	daily := Daily{}
	doc.Find("#tablefix1 > tbody > tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 || i == 1 || i == 2 {
			return
		}
		day := Day{}
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				date, _ := strconv.Atoi(s.Text())
				t := time.Date(targetDate.Year(), targetDate.Month(), date, 0, 0, 0, 0, time.Local)
				day.Date = t
			case 1:
				day.Precipitation.Total = getFloatValue(s.Text())
			case 2:
				day.Precipitation.MaxPrecipitation.Hourly = getFloatValue(s.Text())
			case 3:
				day.Precipitation.MaxPrecipitation.EveryTenMinutes = getFloatValue(s.Text())
			case 4:
				day.Temperature.Average = getFloatValue(s.Text())
			case 5:
				day.Temperature.Highest = getFloatValue(s.Text())
			case 6:
				day.Temperature.Lowest = getFloatValue(s.Text())
			case 7:
				day.Aerovane.AverageWindSpeed = getFloatValue(s.Text())
			case 8:
				day.Aerovane.MaxWindSpeed.Speed = getFloatValue(s.Text())
			case 9:
				day.Aerovane.MaxWindSpeed.Direction = getStringValue(s.Text())
			case 10:
				day.Aerovane.MaxInstantaneousSpeed.Speed = getFloatValue(s.Text())
			case 11:
				day.Aerovane.MaxInstantaneousSpeed.Direction = getStringValue(s.Text())
			case 12:
				day.Aerovane.MostFrequentWindDirection = getStringValue(s.Text())
			case 13:
				day.HoursOfSunlight = getFloatValue(s.Text())
			case 14:
				day.Snow.SnowFall.Total = getFloatValue(s.Text())
			case 15:
				day.Snow.DeepestSnow.Value = getFloatValue(s.Text())
				daily.Days = append(daily.Days, day)
			default:
				return
			}
		})
	})

	return daily, nil
}
