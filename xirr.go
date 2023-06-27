package xirr

import (
	"math"
	"time"
)

func XIRR(payments []Payment) RateInfo {
	var startDate, finishDate = findStartAndFinishDates(payments)
	var totalYears = YearsBetween(startDate, finishDate)
	var yearCashflows = convertToCashflows(payments, startDate)
	var annualRate = calculateXirr(yearCashflows, 0, 1e6, 1e-4)
	var rate = math.Pow(annualRate, totalYears)
	return RateInfo{
		StartDate:  startDate,
		FinishDate: finishDate,
		Years:      totalYears,
		Rate:       rate,
		AnnualRate: annualRate,
	}
}

type cashflow struct {
	Years  float64
	Amount float64
}

func findStartAndFinishDates(payments []Payment) (startDate, finishDate time.Time) {
	var hasValue bool
	for i := range payments {
		var d = payments[i]
		if !hasValue {
			hasValue = true
			startDate = d.Date
			finishDate = d.Date
		} else {
			if d.Date.Before(startDate) {
				startDate = d.Date
			}
			if d.Date.After(finishDate) {
				finishDate = d.Date
			}
		}
	}
	return
}

func convertToCashflows(payments []Payment, baseDate time.Time) []cashflow {
	var result = make([]cashflow, 0, len(payments))
	for _, x := range payments {
		result = append(result, cashflow{
			Years:  YearsBetween(baseDate, x.Date),
			Amount: x.Amount,
		})
	}
	return result
}

func calcEquation(cashflows []cashflow, interestRate float64) float64 {
	var sum float64
	for _, x := range cashflows {
		sum += x.Amount * math.Pow(interestRate, -x.Years)
	}
	return sum
}

func calculateXirr(cashFlows []cashflow, lowRate, highRate, precision float64) float64 {
	var lowResult = calcEquation(cashFlows, lowRate)
	var highResult = calcEquation(cashFlows, highRate)
	for {
		if math.Signbit(lowResult) == math.Signbit(highResult) {
			return math.NaN()
		}

		var middleRate = 0.5 * (lowRate + highRate)
		var middleResult = calcEquation(cashFlows, middleRate)
		if math.Abs(middleResult) <= precision {
			return middleRate
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
