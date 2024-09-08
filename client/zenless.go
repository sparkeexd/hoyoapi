package client

import (
	"github.com/sparkeexd/hoyoapi/internal/components"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/services"
)

// Client that interfaces to HoYoLab endpoints related to Zenless Zone Zero.
// i.e., Daily Reward
type ZenlessClient struct {
	Handler  services.Handler
	Language string
	Daily    components.DailyReward
}

// Constructor.
func NewZenlessClient() *ZenlessClient {
	language := constants.LANG_ENGLISH.String()
	handler := services.NewHandler()

	return &ZenlessClient{
		Handler:  handler,
		Language: language,
		Daily: components.NewDailyReward(
			components.NewDailyRewardParams(constants.ACT_API, constants.ZenlessEventId, constants.ZenlessActId),
			language,
			handler,
		),
	}
}
