package count

import (
	"time"
)

// GetWorkdaysInCurrentMonth returns the number of work days (Monday-Friday)
// in the current month
// NOTE: this will ignore public holidays and leave
func GetWorkdaysInCurrentMonth() int {
	now := time.Now()
	year := now.Year()
	month := now.Month()

	// Get the first day of the current month
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())

	// Get the first day of the next month
	nextMonth := firstDay.AddDate(0, 1, 0)

	workdays := 0

	for current := firstDay; current.Before(nextMonth); current = current.AddDate(0, 0, 1) {
		// Check if the current day is a weekday (Monday-Friday)
		if current.Weekday() != time.Saturday && current.Weekday() != time.Sunday {
			workdays++
		}
	}

	return workdays
}

func GetRemainingWorkDays() int {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	// today := now.Day()

	// Get the first day of the current month
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())

	// Get the first day of the next month
	nextMonth := firstDay.AddDate(0, 1, 0)

	workdays := 0

	for current := now; current.Before(nextMonth); current = current.AddDate(0, 0, 1) {
		// Check if the current day is a weekday (Monday-Friday)
		if current.Weekday() != time.Saturday && current.Weekday() != time.Sunday {
			workdays++
		}
	}

	// exclude today
	return workdays - 1
}

func PercentageToFraction(percentage int) int {
	if percentage < 1 || percentage > 100 {
		return 0
	}
	return 100 / percentage
}
