package task02grasp

type ServicePart struct {
	Name  string
	Price int
}

type PrintSet struct {
	Name  string
	parts []ServicePart
}

func NewPrintSet(name string, parts []ServicePart) PrintSet {
	return PrintSet{Name: name, parts: append([]ServicePart(nil), parts...)}
}

func (p PrintSet) Price() int {
	price := 0
	for _, part := range p.parts {
		price += part.Price
	}
	return price
}
