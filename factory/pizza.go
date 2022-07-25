package main

import "fmt"

type iPizza interface {
	GetName() string
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	iPizza
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (p *Pizza) Prepare() {
	fmt.Println("Preparing", p.Name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Printf("Adding topping:")
	for _, topping := range p.Toppings {
		fmt.Printf(" %v", topping)
	}
	fmt.Println()
}

func (p *Pizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p *Pizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *Pizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *Pizza) GetName() string {
	return p.Name
}

// ---------------NYStyleCheesePizza---------------
type NYStyleCheesePizza struct {
	Pizza
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{
		Pizza: Pizza{
			Name:     "NY Style Sauce and Cheese Pizza",
			Dough:    "Thin Crust Dough",
			Sauce:    "Marinara Sauce",
			Toppings: []string{"Grated Reggiano Cheese"},
		},
	}
}

// ---------------ChicagoStyleCheesePizza---------------
type ChicagoStyleCheesePizza struct {
	Pizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{
		Pizza: Pizza{
			Name:     "Chicago Style Deep Dish Cheese Pizza",
			Dough:    "Extra Thick Crust Dough",
			Sauce:    "Plum Tomato Sauce",
			Toppings: []string{"Shredded Mozzarella Cheese"},
		},
	}
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}
