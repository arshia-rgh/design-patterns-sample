package solid

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int        {}
func (r *Rectangle) SetWidth(width int)   {}
func (r *Rectangle) GetHeight() int       {}
func (r *Rectangle) SetHeight(height int) {}
