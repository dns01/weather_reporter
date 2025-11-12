package main

import (
	"dang/weather_reporter/entity"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const WEATHER_URL = "http://t.weather.itboy.net/api/weather/city/"

//go:embed city.json
var cityJSON []byte

func ListCityCode() error {
	file := cityJSON
	err := json.Unmarshal(file, &cityCode)
	if err != nil {
		panic("json反序列化异常")
	}
	return err
}

func CityCode(name string) (string, error) {
	for _, province := range cityCode.ProvinceList {
		for _, cityInfo := range province.CityList {
			if cityInfo.Name == name {
				return cityInfo.Code, nil
			}
		}
	}
	return "", errors.New("未找到该城市的信息")
}

func weathers(cityCode string) (w entity.WeatherInfo, e error) {
	resp, err := http.Get(WEATHER_URL + cityCode)
	if err != nil {
		panic("http请求异常，请检查服务是否正常")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return w, fmt.Errorf("server returned non-200 status: %d, body: %s", resp.StatusCode, body)
	}

	readAll, e := io.ReadAll(resp.Body)
	if e != nil {
		panic("读取http响应时发生异常")
	}
	var result entity.WeatherInfo
	e = json.Unmarshal(readAll, &result)
	if e != nil {
		panic("json解析异常")
	}
	return result, e

}

func printWeather(info entity.WeatherInfo) {
	if info.Status != 200 {
		fmt.Println("天气数据有误，无法解析")
		return
	}

	f := info.Data.Forecast[0]
	t := info.Data.Forecast[1]

	var strfmt string

	switch info.CityMsg.City {
	case "重庆市", "北京市", "天津市", "上海市":
		strfmt = "%s%s的天气信息：\n"
		info.CityMsg.Parent = ""
	default:
		switch info.CityMsg.Parent {
		case "西藏", "新疆", "广西", "甘肃", "内蒙古":
			strfmt = "%s%s的天气信息：\n"
		default:
			strfmt = "%s省%s的天气信息：\n"
		}
	}

	fmt.Printf(strfmt, info.CityMsg.Parent, info.CityMsg.City)

	fmt.Printf("湿度：%s,  空气质量:%s,  当前温度：%s℃\n活动建议:%s\n\n",
		info.Data.Shidu, info.Data.Quality, info.Data.Wendu, info.Data.Ganmao)
	fmt.Printf("今日天气(%s %s)：%s\n", f.Ymd, f.Week, f.Type)
	fmt.Printf("%s  %s  日出：%s  日落：%s  风向：%s  风速：%s\n", f.High, f.Low, f.Sunrise, f.Sunset, f.Fx, f.Fl)

	fmt.Printf("明日天气(%s %s)：%s\n", t.Ymd, t.Week, t.Type)
	fmt.Printf("%s  %s  日出：%s  日落：%s  风向：%s  风速：%s\n", t.High, t.Low, t.Sunrise, t.Sunset, t.Fx, t.Fl)

}
