package client

import (
	"github.com/sparkeexd/hoyoapi/handler"
	"github.com/sparkeexd/hoyoapi/internal/components"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/middleware"
)

// Client that interfaces to HoYoLab endpoints related to Zenless Zone Zero.
// i.e., Daily Reward
type ZenlessClient struct {
	Handler  *handler.Handler
	Language string
	UserId   int
	Daily    components.DailyReward
}

// Constructor.
func NewZenlessClient(options ClientOptions) *ZenlessClient {
	cookie := middleware.NewCookie(options.ltokenV2, options.ltmidV2, options.ltuidV2)
	handler := handler.NewHandler(cookie)

	return &ZenlessClient{
		Handler:  &handler,
		Language: options.language,
		UserId:   options.userId,
		Daily: components.NewDailyReward(
			components.NewDailyRewardParams(constants.ACT_API, constants.ZenlessEventId, constants.ZenlessActId),
			options.language,
			handler,
		),
	}
}
