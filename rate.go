package xirr

import (
	"fmt"
	"time"
)

type Payment struct {
	Date   time.Time
	Amount float64
}

type RateInfo struct {
	StartDate  time.Time
	FinishDate time.Time
	Years      float64
	Rate       float64
	AnnualRate float64
}

const (
	RateTypeIrr      = 0
	RateTypeArsagera = 1
)

func CalculateRate(payments []Payment, rateType int) RateInfo {
	if rateType == RateTypeIrr {
		return XIRR(payments)
	} else if rateType == RateTypeArsagera {
		return ArsageraRate(payments)
	}
	panic(fmt.Errorf("bad rateType %v", rateType))
}

func YearsBetween(start, finish time.Time) float64 {
	return float64(finish.Sub(start)/(24*time.Hour)) / 365.25
}
