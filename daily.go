package weatherhist

import "time"

const DailyPath = "daily_a1.php"

type Daily struct {
	Days []Day
}

type Day struct {
	Date            time.Time
	Precipitation Precipitation
	Temperature Temperature
	Aerovane Aerovane
	HoursOfSunlight float32
	Snow Snow
}

type Precipitation struct {
	Total            float32
	MaxPrecipitation MaxPrecipitation
}

type MaxPrecipitation struct {
	Hourly          float32
	EveryTenMinutes float32
}

type Temperature struct {
	Average float32
	Highest float32
	Lowest  float32
}

type Aerovane struct {
	AverageWindSpeed          float32
	MaxWindSpeed MaxWindSpeed
	MaxInstantaneousSpeed MaxInstantaneousSpeed
	MostFrequentWindDirection string
}

type MaxWindSpeed struct {
	Speed     float32
	Direction string
}

type MaxInstantaneousSpeed struct {
	Speed     float32
	Direction string
}

type Snow struct {
	SnowFall    SnowFall
	DeepestSnow DeepestSnow
}

type SnowFall struct {
	Total float32
}

type DeepestSnow struct {
	Value float32
}
