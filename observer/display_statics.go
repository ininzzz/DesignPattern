package main

import "fmt"

type StaticsDisplay struct {
	temperatureSum float64
	temperatureMin float64
	temperatureMax float64
	cnt            int32
	weatherData    Subject
}

func NewStaticDisplay(w Subject) *StaticsDisplay {
	d := &StaticsDisplay{
		temperatureMin: 200,
		weatherData:    w,
	}
	d.weatherData.Register(d)
	return d
}

func (d *StaticsDisplay) Update(args ...interface{}) {
	if len(args) < 2 {
		return
	}
	if temperature, ok := args[0].(float64); ok {
		d.temperatureSum += temperature
		if temperature > d.temperatureMax {
			d.temperatureMax = temperature
		}
		if temperature < d.temperatureMin {
			d.temperatureMin = temperature
		}
	}
	d.cnt++
	d.Display()
}

func (d *StaticsDisplay) Name() string {
	return "statics display"
}

func (d *StaticsDisplay) Display() {
	fmt.Println("-------------StaticsDisplay-------------")
	fmt.Printf("Avg/Min/Max:%v/%v/%v\n", (d.temperatureSum / float64(d.cnt)), d.temperatureMin, d.temperatureMax)
}
