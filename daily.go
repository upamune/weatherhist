package weatherhist

import (
	"time"

	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const DailyPath = "daily_%s1.php"

type Daily struct {
	Days []Day `json:"days"`
}

type Day struct {
	Date            time.Time     `json:"date"`
	Precipitation   Precipitation `json:"precipitation"`
	Temperature     Temperature   `json:"temperature"`
	Aerovane        Aerovane      `json:"aerovane"`
	HoursOfSunlight *float32      `json:"hoursOfSunlight"`
	Snow            Snow          `json:"snow"`
	AirPressure     AirPressure   `json:"airPressure"`
	Humidity        Humidity      `json:"humidity"`
}

type Precipitation struct {
	Total            *float32         `json:"total"`
	MaxPrecipitation MaxPrecipitation `json:"maxPrecipitation"`
}

type MaxPrecipitation struct {
	Hourly          *float32 `json:"hourly"`
	EveryTenMinutes *float32 `json:"everyTenMinutes"`
}

type Temperature struct {
	Average *float32 `json:"average"`
	Highest *float32 `json:"highest"`
	Lowest  *float32 `json:"lowest"`
}

type Aerovane struct {
	AverageWindSpeed          *float32              `json:"averageWindSpeed"`
	MaxWindSpeed              MaxWindSpeed          `json:"maxWindSpeed"`
	MaxInstantaneousSpeed     MaxInstantaneousSpeed `json:"maxInstantaneousSpeed"`
	MostFrequentWindDirection *string               `json:"mostFrequentWindDirection"`
}

type MaxWindSpeed struct {
	Speed     *float32 `json:"speed"`
	Direction *string  `json:"direction"`
}

type MaxInstantaneousSpeed struct {
	Speed     *float32 `json:"speed"`
	Direction *string  `json:"direction"`
}

type Snow struct {
	SnowFall    SnowFall    `json:"snowFall"`
	DeepestSnow DeepestSnow `json:"deepestSnow"`
}

type SnowFall struct {
	Total *float32 `json:"total"`
}

type DeepestSnow struct {
	Value *float32 `json:"value"`
}

type AirPressure struct {
	FieldPressure      FieldPressure      `json:"fieldPressure"`
	SeaSurfacePressure SeaSurfacePressure `json:"seaSurfacePressure"`
}

type FieldPressure struct {
	Average *float32 `json:"average"`
}

type SeaSurfacePressure struct {
	Average *float32 `json:"average"`
}

type Humidity struct {
	Average *float32 `json:"average"`
	Lowest  *float32 `json:"lowest"`
}

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = loc
}

func (c *Client) GetDailyData(s Station, targetDate time.Time) (Daily, error) {
	daily, err := c.getDailyDataFromPage(s, targetDate)
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
