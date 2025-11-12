package entity

type WeatherInfo struct {
	Message string  `json:"message"`
	Status  int     `json:"status"`
	Date    string  `json:"date"`
	Time    string  `json:"time"`
	CityMsg CityMsg `json:"cityInfo"`
	Data    Data    `json:"data"`
}

type CityMsg struct {
	City       string `json:"city"`
	Citykey    string `json:"citykey"`
	Parent     string `json:"parent"`
	UpdateTime string `json:"updateTime"`
}

type Forecast struct {
	Date    string `json:"date"`
	High    string `json:"high"`
	Low     string `json:"low"`
	Ymd     string `json:"ymd"`
	Week    string `json:"week"`
	Sunrise string `json:"sunrise"`
	Sunset  string `json:"sunset"`
	Aqi     int    `json:"aqi"`
	Fx      string `json:"fx"`
	Fl      string `json:"fl"`
	Type    string `json:"type"`
	Notice  string `json:"notice"`
}

type Yesterday struct {
	Date    string `json:"date"`
	High    string `json:"high"`
	Low     string `json:"low"`
	Ymd     string `json:"ymd"`
	Week    string `json:"week"`
	Sunrise string `json:"sunrise"`
	Sunset  string `json:"sunset"`
	Aqi     int    `json:"aqi"`
	Fx      string `json:"fx"`
	Fl      string `json:"fl"`
	Type    string `json:"type"`
	Notice  string `json:"notice"`
}

type Data struct {
	Shidu     string     `json:"shidu"`
	Pm25      float64    `json:"pm25"`
	Pm10      float64    `json:"pm10"`
	Quality   string     `json:"quality"`
	Wendu     string     `json:"wendu"`
	Ganmao    string     `json:"ganmao"`
	Forecast  []Forecast `json:"forecast"`
	Yesterday Yesterday  `json:"yesterday"`
}
