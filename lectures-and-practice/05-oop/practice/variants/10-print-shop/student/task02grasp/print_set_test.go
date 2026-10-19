package task02grasp

import "testing"

func sampleProduct() PrintSet {
	return NewPrintSet("Учебный комплект", []ServicePart{
		{Name: "Страницы", Price: 300},
		{Name: "Обложка", Price: 200},
		{Name: "Переплёт", Price: 100},
	})
}

func TestTextReceipt(t *testing.T) {
	got := (Receipt{}).Text(sampleProduct(), 2)
	if got != "Учебный комплект × 2: 1200 руб." {
		t.Fatalf("unexpected receipt: %q", got)
	}
}

func TestHTMLReceipt(t *testing.T) {
	product := sampleProduct()
	product.Name = "Набор <Стандарт> & дополнения"
	got := (Receipt{}).HTML(product, 2)
	if got != "<p>Набор &lt;Стандарт&gt; &amp; дополнения × 2: 1200 руб.</p>" {
		t.Fatalf("unexpected receipt: %q", got)
	}
}

func TestExtendedReceipts(t *testing.T) {
	product := NewPrintSet("Учебный комплект", []ServicePart{
		{Name: "Страницы", Price: 300},
		{Name: "Обложка", Price: 200},
		{Name: "Переплёт", Price: 100},
		{Name: "Ламинация", Price: 150},
	})
	receipt := Receipt{}
	if got := receipt.Text(product, 2); got != "Учебный комплект × 2: 1500 руб." {
		t.Fatalf("unexpected text receipt: %q", got)
	}
	if got := receipt.HTML(product, 2); got != "<p>Учебный комплект × 2: 1500 руб.</p>" {
		t.Fatalf("unexpected HTML receipt: %q", got)
	}
}

func TestPrintSetCopiesParts(t *testing.T) {
	parts := []ServicePart{{Name: "Страницы", Price: 300}}
	product := NewPrintSet("Комплект", parts)
	parts[0].Price = 999
	if got := product.Price(); got != 300 {
		t.Fatalf("input slice changed price: got %d, want 300", got)
	}
}
