package solid

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	size  Size
	color Color
}

/*
	Here if we want to add new filter or combine of the filters we need to repeatedly modify the existing Filter struct
	and this will break the open-closed principle
*/

// Filter actually breaks the OCP because it will be modify over and over
type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {}

// More ... ( we need modify the Filter type) --> breaks the OCP

// -- (Solution) --
// We can use `Specification pattern`

type Specification interface {
	IsSatisfied(p *Product) bool
}

/*
	We can implement types for each Filter we want to add and
	made them all implement the same interface
*/

type ColorSpecification struct {
	color Color
}

func (cs *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == cs.color
}
