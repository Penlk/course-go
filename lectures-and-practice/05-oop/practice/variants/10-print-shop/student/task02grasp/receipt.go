package task02grasp

import (
	"fmt"
	"html"
)

type Receipt struct{}

func (Receipt) Text(product PrintSet, quantity int) string {
	return fmt.Sprintf("%s × %d: %d руб.", product.Name, quantity, product.Price()*quantity)
}

func (Receipt) HTML(product PrintSet, quantity int) string {
	return fmt.Sprintf("<p>%s × %d: %d руб.</p>", html.EscapeString(product.Name), quantity, product.Price()*quantity)
}
