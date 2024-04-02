package decorator

type Beverage interface {
	Cost() int
	Description() string
}

type Espresso struct{}

func (e *Espresso) Cost() int {
	return 2
}

func (e *Espresso) Description() string {
	return "Espresso"
}

func WithMilk(beverage Beverage) Beverage {
	return &WithMilkDecorator{beverage}
}

type WithMilkDecorator struct {
	beverage Beverage
}

func (w *WithMilkDecorator) Cost() int {
	return w.beverage.Cost() + 1
}

func (w *WithMilkDecorator) Description() string {
	return w.beverage.Description() + ", with Milk"
}
