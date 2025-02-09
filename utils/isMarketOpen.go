package utils

import (
	"fmt"
	"nepseserver/server"
	"time"
)

type contextKey string

const MarketStatusKey contextKey = "market_status"

func isMarketOpen() (bool, error) {
	loc, err := time.LoadLocation("Asia/Kathmandu")
	if err != nil {
		return false, err
	}
	now := time.Now().In(loc)

	// Check if today is Friday (5) or Saturday (6) - market closed
	weekday := now.Weekday()
	if weekday == time.Friday || weekday == time.Saturday {
		fmt.Println("Market is CLOSED today .")

		return false, fmt.Errorf("market is CLOSED ")
	}

	// Check market open time (11 AM to 3 PM)
	currentHour := now.Hour()
	if currentHour < 11 || currentHour >= 15 {

		return false, fmt.Errorf("market is CLOSED ")
	}

	// Wait for 2 minutes before fetching NEPSE data
	fmt.Println("Waiting for 2 minutes before fetching NEPSE data...")
	time.Sleep(3 * time.Minute)

	// Fetch NEPSE data and check AsOfDate
	nepseIndex, err := server.FetchNepseData()
	if err != nil {
		return false, err
	}

	// Get current date in "YYYY-MM-DD" format
	currentDate := now.Format("2006-01-02")

	// Check if AsOfDate matches the current date
	if nepseIndex.AsOfDate != currentDate {
		return false, fmt.Errorf("market is Closed")
	}

	return true, nil
}

var IsMarketOpenGlobal bool

func Init() {
	open, err := isMarketOpen()
	if err != nil {
		IsMarketOpenGlobal = false
	} else {
		IsMarketOpenGlobal = open
	}
}
