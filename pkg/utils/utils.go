package utils

import (
	"fmt"
	"time"
)

// GetMonthVal - returns Month format of a int month
func GetMonthVal(m int) time.Month {
	months := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
	return months[m-1]
}

// CheckErr
func CheckErr(e error) bool {
	if e != nil {
		fmt.Println(e)
		return true
	}
	return false
}
