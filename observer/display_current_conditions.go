package main

import "fmt"

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

func NewCurrentConditionsDisplay(w Subject) *CurrentConditionsDisplay {
	d := &CurrentConditionsDisplay{
		weatherData: w,
	}
	d.weatherData.Register(d)
	return d
}

func (d *CurrentConditionsDisplay) Update(args ...interface{}) {
	if len(args) < 2 {
		return
	}
	var ok bool
	if d.temperature, ok = args[0].(float64); !ok {
		return
	}
	if d.humidity, ok = args[1].(float64); !ok {
		return
	}
	d.Display()
}

func (d *CurrentConditionsDisplay) Name() string {
	return "current conditions display"
}

func (d *CurrentConditionsDisplay) Display() {
	fmt.Println("--------CurrentConditionsDisplay--------")
	fmt.Printf("temperature: %v\n", d.temperature)
	fmt.Printf("humidity: %v\n", d.humidity)
}
