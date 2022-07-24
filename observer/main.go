package main

import "fmt"

func main() {
	wd := NewWeatherData()

	_ = NewCurrentConditionsDisplay(wd)
	_ = NewStaticDisplay(wd)

	wd.SetMeasurements(80.7, 65, 30.4)
	fmt.Println()
	wd.SetMeasurements(82.3, 70, 29.2)
}
