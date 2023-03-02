package product

type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shirt struct {
	logo string
	size int
}

func NewShirt(l string, s int) *Shirt {
	return &Shirt{
		logo: l,
		size: s,
	}
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}
