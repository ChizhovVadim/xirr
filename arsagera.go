package xirr

import (
	"math"
	"sort"
)

// https://www.youtube.com/watch?v=9_Lj1CSbAh0&t=497s
func ArsageraRate(payments []Payment) RateInfo {
	payments = clonePayments(payments)
	sort.Slice(payments, func(i, j int) bool {
		return payments[i].Date.Before(payments[j].Date)
	})

	var minDate = payments[0].Date
	var maxDate = payments[len(payments)-1].Date
	var totalYears = YearsBetween(minDate, maxDate)

	var workingAmount float64
	var workingSum float64
	var totalPnL float64
	for i := range payments {
		var current = payments[i].Amount
		totalPnL += current
		if i < len(payments)-1 {
			workingSum -= current
			if workingSum > 0 {
				var weight = YearsBetween(payments[i].Date, payments[i+1].Date) / totalYears
				workingAmount += workingSum * weight
			}
		}
	}

	var rate = math.Max(0, 1+totalPnL/workingAmount)
	var annualRate = math.Pow(rate, 1.0/totalYears)

	return RateInfo{
		StartDate:  minDate,
		FinishDate: maxDate,
		Years:      totalYears,
		Rate:       rate,
		AnnualRate: annualRate,
	}
}

func clonePayments(payments []Payment) []Payment {
	var result = make([]Payment, len(payments))
	copy(result, payments)
	return result
}
