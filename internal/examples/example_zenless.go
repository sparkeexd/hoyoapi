package main

import (
	"github.com/sparkeexd/hoyoapi/hoyoapi"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
)

// Claim Zenless daily rewards.
func ZenlessDailyReward(options hoyoapi.ClientOptions) {
	zenless := hoyoapi.NewZenlessClient(options)
	response, err := zenless.Daily.Claim()
	utilities.PrintJSON(response, err)
}
