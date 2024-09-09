package client

import (
	"github.com/sparkeexd/hoyoapi/internal/components"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/services"
)

// Client that interfaces to HoYoLab endpoints related to Star Rail.
// i.e., Daily Reward
type StarRailClient struct {
	Handler  services.Handler
	Language string
	Daily    components.DailyReward
}

// Constructor.
func NewStarRailClient() *StarRailClient {
	language := constants.LANG_ENGLISH.String()
	handler := services.NewHandler()

	return &StarRailClient{
		Handler:  handler,
		Language: language,
		Daily: components.NewDailyReward(
			components.NewDailyRewardParams(constants.PUBLIC_API, constants.StarRailEventID, constants.StarRailActID),
			language,
			handler,
		),
	}
}
