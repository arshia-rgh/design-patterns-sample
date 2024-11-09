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

type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {}

// More ... ( we need modify the Filter type) --> breaks the OCP
