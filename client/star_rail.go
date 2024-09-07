package client

import (
	"github.com/sparkeexd/hoyoapi/handlers"
	"github.com/sparkeexd/hoyoapi/internal/components"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/middleware"
)

// Client that interfaces to HoYoLab endpoints related to Star Rail.
// i.e., Daily Reward
type StarRailClient struct {
	Cache    *middleware.Cache
	Handler  *handlers.Handler
	Language string
	UserId   int
	Daily    components.DailyReward
}

// Constructor.
func NewStarRailClient(options ClientOptions) *StarRailClient {
	cookie := middleware.NewCookie(options.ltokenV2, options.ltmidV2, options.ltuidV2)
	handler := handlers.NewHandler(cookie)

	return &StarRailClient{
		Handler:  &handler,
		Cache:    middleware.NewCache(),
		Language: options.language,
		UserId:   options.userId,
		Daily: components.NewDailyReward(
			components.NewDailyRewardParams(constants.PUBLIC_API, constants.StarRailEventId, constants.StarRailActId),
			options.language,
			&handler,
		),
	}
}
