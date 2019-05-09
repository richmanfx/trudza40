package test

import (
	"testing"
	"trudza40/controllers"
)

func Test_GetRentalPeriodFromString(t *testing.T) {

	type testPair struct {
		rentPeriodString string // Срок действия договора - строка со страницы (лет, мес.)
		rentPeriodInt    int    // Срок действия договора, месяцы
	}

	var testsData = []testPair{
		{"15 лет", 180},
		{"5 лет", 60},
		{"3 г.", 36},
		{"11 мес.", 11},
	}

	for _, pair := range testsData {

		actualResult := controllers.GetRentalPeriodFromString(pair.rentPeriodString)
		if actualResult != pair.rentPeriodInt {
			t.Errorf("Для '%s' ожидался результат '%d', а получен '%v'",
				pair.rentPeriodString, pair.rentPeriodInt, actualResult)
		}
	}

}
