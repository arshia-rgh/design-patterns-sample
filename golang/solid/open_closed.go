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

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

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

// if we need new filter or anything we can just add new type
// and not modify the existing one

type SizeSpecification struct {
	size Size
}

func (ss *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == ss.size
}

// For combined filters like size and color we can use composite pattern

type AndSpecification struct {
	first, second Specification
}

func (as *AndSpecification) IsSatisfied(p *Product) bool {
	return as.first.IsSatisfied(p) && as.second.IsSatisfied(p)
}

/*
	This filter type now doesnt break the OCP
*/

type BetterFilter struct{}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}
