package task03interface

import (
	"errors"
	"fmt"

	"example.com/oop-practice/10-print-shop/student/task03interface/clients"
)

type OfficeAdapter struct {
	Client *clients.OfficePrinter
}

func (a OfficeAdapter) Estimate(format string, pages int) (PrintQuote, error) {
	price, minutes, err := a.Client.Estimate(format, pages)
	if errors.Is(err, clients.ErrOfficeFormatUnavailable) {
		return PrintQuote{}, fmt.Errorf("OfficePrinter: %w", ErrFormatUnavailable)
	}
	if err != nil {
		return PrintQuote{}, err
	}
	return PrintQuote{PriceRubles: price, Minutes: minutes}, nil
}

type QuickAdapter struct {
	Client *clients.QuickPrinter
}

func (a QuickAdapter) Estimate(format string, pages int) (PrintQuote, error) {
	result, err := a.Client.Calculate(clients.PrintRequest{
		PaperFormat: format,
		PageCount:   pages,
	})
	if errors.Is(err, clients.ErrQuickFormatUnavailable) {
		return PrintQuote{}, fmt.Errorf("QuickPrinter: %w", ErrFormatUnavailable)
	}
	if err != nil {
		return PrintQuote{}, err
	}
	return PrintQuote{
		PriceRubles: roundUp(result.CostKopecks, 100),
		Minutes:     roundUp(result.DurationSeconds, 60),
	}, nil
}

func roundUp(value, unit int) int {
	result := value / unit
	if value%unit != 0 {
		result++
	}
	return result
}
