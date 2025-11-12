package entity

type CityInfo struct {
	Name string `json:"市名"`
	Code string `json:"编码"`
}

type Province struct {
	Name     string     `json:"省"`
	CityList []CityInfo `json:"市"`
}

type CityCode struct {
	ProvinceList []Province `json:"城市代码"`
}
