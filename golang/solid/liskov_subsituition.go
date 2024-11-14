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

func NewSquare(size int) *Square {
	square := Square{}
	square.height = size
	square.width = size
	return &square
}

// we should set both high and width to keep this square

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// here in UseIt we will have wrong result with Square and we already violates the liskov here

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetHeight() * sized.GetWidth()

	fmt.Println("expected: ", expectedArea, "actual: ", actualArea)
}

// solution is to do like this

type Square2 struct {
	size int // width, height
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

// having rectangle and square together with the same interface
//will break the liskov substitution and should be separated
