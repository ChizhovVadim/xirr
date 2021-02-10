package xirr

import (
	"math"
	"testing"
	"time"
)

func TestXirr(t *testing.T) {
	var tests = []struct {
		Payments []Payment
		Xirr     float64
	}{
		{
			Payments: []Payment{
				Payment{mustParseDate("2018-01-01"), -500000},
				Payment{mustParseDate("2018-08-22"), 500000},
				Payment{mustParseDate("2020-03-13"), 1000000},
				Payment{mustParseDate("2020-03-18"), 400000},
				Payment{mustParseDate("2020-04-16"), 631928.98},
				Payment{mustParseDate("2021-01-05"), -1000000},
				Payment{mustParseDate("2021-02-10"), 794000},
			},
			Xirr: 2.591,
		},
	}
	for i, test := range tests {
		var xirr, err = XIRR(test.Payments)
		if err != nil {
			t.Error(err)
			continue
		}
		if math.Abs(xirr-test.Xirr) > 0.005 {
			t.Errorf("%v %v %v", i, test, xirr)
			continue
		}
	}
}

func mustParseDate(s string) time.Time {
	var t, err = time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t
}
