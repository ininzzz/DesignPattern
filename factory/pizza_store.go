package main

type iPizzaStore interface {
	CreatePizza(t string) iPizza
	OrderPizza(t string) iPizza
}

type PizzaStore struct {
	iPizzaStore
}

func (s *PizzaStore) OrderPizza(t string) iPizza {
	pizza := s.CreatePizza(t)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

// ---------------NYPizzaStore---------------
type NYPizzaStore struct {
	PizzaStore
}

func NewNYPizzaStore() *NYPizzaStore {
	return &NYPizzaStore{
		PizzaStore: PizzaStore{
			iPizzaStore: &NYPizzaStore{},
		},
	}
}

func (s *NYPizzaStore) CreatePizza(t string) iPizza {
	switch t {
	case "cheese":
		return NewNYStyleCheesePizza()
	default:
		return nil
	}
}

// ---------------ChicagoPizzaStore---------------
type ChicagoPizzaStore struct {
	PizzaStore
}

func NewChicagoPizzaStore() *ChicagoPizzaStore {
	return &ChicagoPizzaStore{
		PizzaStore: PizzaStore{
			iPizzaStore: &ChicagoPizzaStore{},
		},
	}
}

func (s *ChicagoPizzaStore) CreatePizza(t string) iPizza {
	switch t {
	case "cheese":
		return NewChicagoStyleCheesePizza()
	default:
		return nil
	}
}
