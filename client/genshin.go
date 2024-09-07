package client

import (
	"fmt"
	"net/http"

	"github.com/sparkeexd/hoyoapi/handlers"
	"github.com/sparkeexd/hoyoapi/internal/components"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/internal/errors"
	"github.com/sparkeexd/hoyoapi/middleware"
)

// Client that interfaces to HoYoLab endpoints related to Genshin Impact.
// i.e., Spiral Abyss, Daily Reward
type GenshinClient struct {
	Cache    *middleware.Cache
	Handler  *handlers.Handler
	Language string
	UserId   int
	Daily    components.DailyReward
}

// Constructor.
func NewGenshinClient(options ClientOptions) GenshinClient {
	cookie := middleware.NewCookie(options.ltokenV2, options.ltmidV2, options.ltuidV2)
	handler := handlers.NewHandler(cookie)

	return GenshinClient{
		Handler:  &handler,
		Cache:    middleware.NewCache(),
		Language: options.language,
		UserId:   options.userId,
		Daily: components.NewDailyReward(
			components.NewDailyRewardParams(constants.HK4E_API, constants.GenshinEventId, constants.GenshinActId),
			options.language,
			&handler,
		),
	}
}

// Get current Spiral Abyss information.
// Set current argument to false to get previous cycle's information.
func (genshin GenshinClient) SpiralAbyss(current bool) (map[string]interface{}, error) {
	scheduleType := 1
	if !current {
		scheduleType = 2
	}

	server, err := genshin.getRegion()
	if err != nil {
		return nil, err
	}

	request := handlers.NewRequest(constants.GENSHIN_RECORD_SPIRAL_ABYSS_API, http.MethodGet).
		AddParam("role_id", fmt.Sprint(genshin.UserId)).
		AddParam("server", server.String()).
		AddParam("schedule_type", fmt.Sprint(scheduleType)).
		AddDynamicSecret(constants.DS_GLOBAL).
		AddLanguage(genshin.Language).
		Build()

	data := make(map[string]interface{})
	err = genshin.Handler.Send(request, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Get Genshin regional server code based off the user ID.
func (genshin GenshinClient) getRegion() (*constants.GenshinRegion, error) {
	// Get 1st digit of user ID to determine the region.
	i := genshin.UserId
	for i >= 10 {
		i /= 10
	}

	region, exists := constants.GenshinServerRegions[i]
	if !exists {
		return nil,
			errors.NewError(
				errors.REGION_SERVER_CODE_ERROR,
				fmt.Sprintf("Genshin regional server code not found for UID %d", genshin.UserId),
			)
	}

	return &region, nil
}
