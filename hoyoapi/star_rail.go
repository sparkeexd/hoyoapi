package hoyoapi

import (
	"github.com/sparkeexd/hoyoapi/internal/components"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/internal/handler"
	"github.com/sparkeexd/hoyoapi/internal/middleware"
)

// Client that interfaces to HoYoLab endpoints related to Star Rail.
// i.e., Daily Reward
type StarRailClient struct {
	Handler  *handler.Handler
	Language string
	UserId   int
	Daily    components.DailyReward
}

// Constructor.
func NewStarRailClient(options ClientOptions) *StarRailClient {
	cookie := middleware.NewCookie(options.ltokenV2, options.ltmidV2, options.ltuidV2)
	handler := handler.NewHandler(cookie)

	return &StarRailClient{
		Handler:  &handler,
		Language: options.language,
		UserId:   options.userId,
		Daily: components.NewDailyReward(
			constants.GAME_STAR_RAIL,
			options.language,
			&handler,
		),
	}
}
