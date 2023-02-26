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
				{mustParseDate("2018-01-01"), -500000},
				{mustParseDate("2018-08-22"), 500000},
				{mustParseDate("2020-03-13"), 1000000},
				{mustParseDate("2020-03-18"), 400000},
				{mustParseDate("2020-04-16"), 631928.98},
				{mustParseDate("2021-01-05"), -1000000},
				{mustParseDate("2021-02-10"), 794000},
			},
			Xirr: 2.591,
		},
		{
			Payments: []Payment{
				Payment{mustParseDate("2022-08-26"), -20},
				Payment{mustParseDate("2022-09-26"), 0.12},
				Payment{mustParseDate("2022-10-26"), 0.12},
				Payment{mustParseDate("2022-11-28"), 0.12},
				Payment{mustParseDate("2022-12-26"), 0.12},
				Payment{mustParseDate("2023-01-26"), 20.12},
			},
			Xirr: 1.074,
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
