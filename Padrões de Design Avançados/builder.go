package builder

type Product struct {
	Part1 string
	Part2 int
}

type Builder interface {
	BuildPart1() string
	BuildPart2() int
}

type ConcreteBuilder struct{}

func (c *ConcreteBuilder) BuildPart1() string {
	return "Part1"
}

func (c *ConcreteBuilder) BuildPart2() int {
	return 42
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder}
}

func (d *Director) Construct() *Product {
	return &Product{
		Part1: d.builder.BuildPart1(),
		Part2: d.builder.BuildPart2(),
	}
}
