package task03interface

import "errors"

var ErrFormatUnavailable = errors.New("print format unavailable")

type PrintQuote struct {
	PriceRubles int
	Minutes     int
}

type PrintEstimator interface {
	Estimate(format string, pages int) (PrintQuote, error)
}

type Checkout struct {
	client PrintEstimator
}

func NewCheckout(client PrintEstimator) *Checkout {
	return &Checkout{client: client}
}

func (s *Checkout) QuotePrint(format string, pages int) (PrintQuote, error) {
	quote, err := s.client.Estimate(format, pages)
	if err != nil {
		return PrintQuote{}, err
	}
	return quote, nil
}
