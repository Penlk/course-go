package task03interface

import (
	"errors"
	"testing"

	"example.com/oop-practice/10-print-shop/student/task03interface/clients"
)

type estimatorStub func(string, int) (PrintQuote, error)

func (s estimatorStub) Estimate(format string, pages int) (PrintQuote, error) {
	return s(format, pages)
}

func TestCheckoutWithStub(t *testing.T) {
	failure := errors.New("printer offline")
	for _, err := range []error{nil, failure, ErrFormatUnavailable} {
		calls := 0
		want := PrintQuote{PriceRubles: 42, Minutes: 7}
		service := NewCheckout(estimatorStub(func(format string, pages int) (PrintQuote, error) {
			calls++
			if format != "A3" || pages != 12 {
				t.Fatalf("unexpected request: %q, %d", format, pages)
			}
			return want, err
		}))
		got, gotErr := service.QuotePrint("A3", 12)
		if calls != 1 || !errors.Is(gotErr, err) {
			t.Fatalf("calls=%d, error=%v, want %v", calls, gotErr, err)
		}
		if err != nil {
			want = PrintQuote{}
		}
		if got != want {
			t.Fatalf("quote=%+v, want %+v", got, want)
		}
	}
}

func TestEstimatorContract(t *testing.T) {
	providers := []struct {
		name      string
		estimator PrintEstimator
		want      PrintQuote
	}{
		{"office", OfficeAdapter{Client: &clients.OfficePrinter{}}, PrintQuote{350, 15}},
		{"quick", QuickAdapter{Client: &clients.QuickPrinter{}}, PrintQuote{351, 10}},
	}
	for _, provider := range providers {
		t.Run(provider.name, func(t *testing.T) {
			service := NewCheckout(provider.estimator)
			for _, format := range []string{"A4", "A3"} {
				t.Run(format, func(t *testing.T) {
					got, err := service.QuotePrint(format, 25)
					if err != nil || got != provider.want {
						t.Fatalf("quote=%+v, error=%v, want %+v", got, err, provider.want)
					}
				})
			}
			for _, tc := range []struct {
				name        string
				format      string
				pages       int
				unavailable bool
			}{
				{"unsupported", "A5", 25, true},
				{"unknown", "custom", 25, true},
				{"empty", "", 25, false},
				{"zero", "A4", 0, false},
				{"negative", "A4", -1, false},
			} {
				t.Run(tc.name, func(t *testing.T) {
					got, err := service.QuotePrint(tc.format, tc.pages)
					if err == nil || errors.Is(err, ErrFormatUnavailable) != tc.unavailable || got != (PrintQuote{}) {
						t.Fatalf("quote=%+v, error=%v, want unavailable=%v", got, err, tc.unavailable)
					}
				})
			}
		})
	}
}
