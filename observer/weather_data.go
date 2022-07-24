package main

type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make([]Observer, 0),
	}
}

func (w *WeatherData) Register(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) Remove(o Observer) {
	for i, observer := range w.observers {
		if observer.Name() == o.Name() {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			return
		}
	}
}

func (w *WeatherData) Notify() {
	for _, observer := range w.observers {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) Changed() {
	w.Notify()
}

func (w *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.Changed()
}
