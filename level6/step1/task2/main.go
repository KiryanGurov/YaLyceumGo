package main

type WeatherCondition int

const (
	Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	var conditionMultiplier = map[WeatherCondition]float64{
		Rain:      0.125,
		HeavyRain: 0.2,
		Snow:      0.15,
	}
	
	res := 1.0 + conditionMultiplier[weather.Condition]
	if weather.WindSpeed > 15 {
		return res + 0.1
	}
	return res
}
