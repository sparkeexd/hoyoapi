package main

import (
	"github.com/sparkeexd/hoyoapi/client"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
	"github.com/sparkeexd/hoyoapi/services"
)

// Claim Star Rail daily rewards.
func StarRailDailyReward(cookie services.Cookie) {
	starRail := client.NewStarRailClient()
	response, err := starRail.Daily.Claim(cookie)
	utilities.PrintJSON(response, err)
}
