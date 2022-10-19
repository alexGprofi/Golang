package main

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (sc TemperatureUnit) String() string {
	units := []string{"ºC", "ºF"}
	return units[sc]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (sc SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[sc]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (met MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", met.location, met.temperature, met.windDirection, met.windSpeed, met.humidity)
}

func main() {
	//temperatureUnit := Celsius
	celsiusUnit := Celsius
	fahrenheitUnit := Fahrenheit

	fmt.Println(celsiusUnit.String())
	// => °C
	fmt.Println(fahrenheitUnit.String())
	// => °F
	fmt.Println(fmt.Sprint(celsiusUnit))
	// Output: °C

	celsiusTemp := Temperature{
		degree: 21,
		unit:   Celsius,
	}
	fmt.Println(celsiusTemp.String())
	// => 21 °C
	fmt.Println(fmt.Sprint(celsiusTemp))
	// Output: 21 °C

	fahrenheitTemp := Temperature{
		degree: 75,
		unit:   Fahrenheit,
	}
	fmt.Println(fahrenheitTemp.String())
	// => 75 °F
	fmt.Println(fmt.Sprint(fahrenheitTemp))
	// Output: 75 °F
	mphUnit := MilesPerHour
	fmt.Println(mphUnit.String())
	// => mph
	fmt.Println(fmt.Sprint(mphUnit))
	// Output: mph

	kmhUnit := KmPerHour
	fmt.Println(kmhUnit.String())
	// => km/h
	fmt.Println(fmt.Sprint(kmhUnit))
	// Output: km/h
	windSpeedNow := Speed{
		magnitude: 18,
		unit:      KmPerHour,
	}
	fmt.Println(windSpeedNow.String())
	// => 18 km/h
	fmt.Println(windSpeedNow)
	// Output: 18 km/h

	windSpeedYesterday := Speed{
		magnitude: 22,
		unit:      MilesPerHour,
	}
	fmt.Println(windSpeedYesterday.String())
	// => 22 mph
	fmt.Println(fmt.Sprint(windSpeedYesterday))
	// Output: 22 mph

	sfData := MeteorologyData{
		location: "San Francisco",
		temperature: Temperature{
			degree: 57,
			unit:   Fahrenheit,
		},
		windDirection: "NW",
		windSpeed: Speed{
			magnitude: 19,
			unit:      MilesPerHour,
		},
		humidity: 60,
	}

	fmt.Println(sfData.String())
	// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
}
