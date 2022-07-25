package main

import "fmt"

func main() {
	nyStore := NewNYPizzaStore()
	chicagoStore := NewChicagoPizzaStore()

	pizza := nyStore.OrderPizza("cheese")
	fmt.Println(pizza.GetName(), "ordered")
	pizza = chicagoStore.OrderPizza("cheese")
	fmt.Println(pizza.GetName(), "ordered")
}
