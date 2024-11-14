package solid

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int        { return r.width }
func (r *Rectangle) SetWidth(width int)   { r.width = width }
func (r *Rectangle) GetHeight() int       { return r.height }
func (r *Rectangle) SetHeight(height int) { r.height = height }

type Square struct {
	Rectangle
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetHeight() * sized.GetWidth()

	fmt.Println("expected: ", expectedArea, "actual: ", actualArea)
}
