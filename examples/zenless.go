package main

import (
	"github.com/sparkeexd/hoyoapi/client"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
	"github.com/sparkeexd/hoyoapi/services"
)

// Claim Zenless daily rewards.
func ZenlessDailyReward(cookie services.Cookie) {
	zenless := client.NewZenlessClient()
	response, err := zenless.Daily.Claim(cookie)
	utilities.PrintJSON(response, err)
}
