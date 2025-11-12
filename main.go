package main

import (
	"dang/weather_reporter/entity"
	"fmt"
	"log"
	"os"
)

var cityCode entity.CityCode

func main() {
	var cityName = "郑州"
	// 读取json文件，将其转化为结构体
	err := ListCityCode()
	if err != nil {
		log.Fatal("读取json文件发生异常")
	}

	if len(os.Args) < 2 {
		fmt.Println("未指定城市，使用默认城市：郑州")
	} else {
		cityName = os.Args[1]
	}
	// 根据城市名称，获取城市编码
	cityNumber, err := CityCode(cityName)
	if err != nil {
		log.Fatal("获取城市编码失败")
	}

	// 使用完整的url，获取天气信息
	weatherInfo, err := weathers(cityNumber)
	if err != nil {
		log.Fatal("请求天气数据失败")
	}

	// 打印天气信息
	printWeather(weatherInfo)

}
