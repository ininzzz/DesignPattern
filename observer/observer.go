package main

type Observer interface {
	Update(...interface{})
	Name() string
}

type Subject interface {
	Register(o Observer)
	Remove(o Observer)
	Notify()
}
