package xirr

import (
	"math"
	"sort"
)

// alternative way to calculate rate
// https://www.youtube.com/watch?v=9_Lj1CSbAh0&t=497s
func ArsageraRate(payments []Payment) (float64, error) {
	payments = clonePayments(payments)
	sort.Slice(payments, func(i, j int) bool {
		return payments[i].Date.Before(payments[j].Date)
	})
	var minDate = payments[0].Date
	var maxDate = payments[len(payments)-1].Date
	var totalYears = yearsBetween(minDate, maxDate)

	var workingAmount float64
	var workingSum float64
	var totalPnL float64
	for i := range payments {
		var current = payments[i].Amount
		totalPnL += current
		if i < len(payments)-1 {
			workingSum -= current
			if workingSum > 0 {
				var weight = yearsBetween(payments[i].Date, payments[i+1].Date) / totalYears
				workingAmount += workingSum * weight
			}
		}
	}

	var rate = math.Pow(1+totalPnL/workingAmount, 1.0/totalYears)

	return rate, nil
}

func clonePayments(payments []Payment) []Payment {
	var result = make([]Payment, len(payments))
	copy(result, payments)
	return result
}
