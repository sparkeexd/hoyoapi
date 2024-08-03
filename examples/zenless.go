package main

import (
	"github.com/sparkeexd/hoyoapi/client"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
)

// Claim Zenless daily rewards.
func ZenlessDailyReward(options client.ClientOptions) {
	zenless := client.NewZenlessClient(options)
	response, err := zenless.Daily.Claim()
	utilities.PrintJSON(response, err)
}
