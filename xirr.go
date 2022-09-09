package xirr

import (
	"errors"
	"math"
	"time"
)

var errXirr = errors.New("xirr cannot be calculated")

type Payment struct {
	Date   time.Time
	Amount float64
}

type cashflow struct {
	Years  float64
	Amount float64
}

func XIRR(payments []Payment) (float64, error) {
	var yearCashflows = convertToCashflows(payments)
	const decimals = 3
	var precision = math.Pow(10, -decimals)
	var minRate = precision
	const maxRate = 1000000
	return calculateXirr(yearCashflows, minRate, maxRate, precision, decimals)
}

func findMinDate(payments []Payment) time.Time {
	var result = payments[0].Date
	for i := 1; i < len(payments); i++ {
		var d = payments[i].Date
		if d.Before(result) {
			result = d
		}
	}
	return result
}

func convertToCashflows(payments []Payment) []cashflow {
	var minDate = findMinDate(payments)
	var result = make([]cashflow, len(payments))
	for i, x := range payments {
		result[i] = cashflow{
			Years:  yearsBetween(minDate, x.Date),
			Amount: x.Amount,
		}
	}
	return result
}

func yearsBetween(start, finish time.Time) float64 {
	return float64(finish.Sub(start)/(24*time.Hour)) / 365.25
}

func calcEquation(cashflows []cashflow, interestRate float64) float64 {
	var sum float64
	for _, x := range cashflows {
		sum += x.Amount * math.Pow(interestRate, -x.Years)
	}
	return sum
}

func calculateXirr(cashFlows []cashflow, lowRate, highRate, precision float64, decimals int) (float64, error) {
	var lowResult = calcEquation(cashFlows, lowRate)
	var highResult = calcEquation(cashFlows, highRate)
	for {
		if math.Signbit(lowResult) == math.Signbit(highResult) {
			return 0, errXirr
		}

		var middleRate = 0.5 * (lowRate + highRate)
		var middleResult = calcEquation(cashFlows, middleRate)
		if math.Abs(middleResult) <= precision {
			return middleRate, nil
		}
		if math.Signbit(middleResult) == math.Signbit(lowResult) {
			lowRate = middleRate
			lowResult = middleResult
		} else {
			highRate = middleRate
			highResult = middleResult
		}
	}
}
