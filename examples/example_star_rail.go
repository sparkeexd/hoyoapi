package main

import (
	"github.com/sparkeexd/hoyoapi/client"
	"github.com/sparkeexd/hoyoapi/utilities"
)

// Claim Star Rail daily rewards.
func StarRailDailyReward(options client.ClientOptions) {
	starRail := client.NewStarRailClient(options)
	response, err := starRail.Daily.Claim()
	utilities.PrintJSON(response, err)
}
