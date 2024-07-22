package main

import (
	"github.com/sparkeexd/hoyoapi/hoyoapi"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
)

// Claim Star Rail daily rewards.
func StarRailDailyReward(options hoyoapi.ClientOptions) {
	starRail := hoyoapi.NewStarRailClient(options)
	response, err := starRail.Daily.Claim()
	utilities.PrintJSON(response, err)
}
