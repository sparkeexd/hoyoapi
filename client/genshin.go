package client

import (
	"fmt"
	"net/http"

	"github.com/sparkeexd/hoyoapi/internal/components"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/internal/errors"
	"github.com/sparkeexd/hoyoapi/models"
	"github.com/sparkeexd/hoyoapi/services"
)

// Client that interfaces to HoYoLab endpoints related to Genshin Impact.
// i.e., Spiral Abyss, Daily Reward
type GenshinClient struct {
	Language string
	Daily    components.DailyReward
	handler  services.Handler
}

// Constructor.
func NewGenshinClient() *GenshinClient {
	language := constants.LANG_ENGLISH.String()
	handler := services.NewHandler()

	return &GenshinClient{
		handler:  handler,
		Language: language,
		Daily: components.NewDailyReward(
			components.NewDailyRewardParams(constants.HK4E_API, constants.GenshinEventID, constants.GenshinActID),
			language,
			handler,
		),
	}
}

// Get current Spiral Abyss information.
// Set current argument to false to get previous cycle's information.
func (genshin GenshinClient) SpiralAbyss(cookie services.Cookie, userID int, current bool) (models.SpiralAbyss, error) {
	var res models.SpiralAbyss

	scheduleType := 1
	if !current {
		scheduleType = 2
	}

	server, err := genshin.getRegion(userID)
	if err != nil {
		return res, err
	}

	request := services.NewRequest(constants.GENSHIN_RECORD_SPIRAL_ABYSS_API, http.MethodGet, cookie).
		AddParam("role_id", fmt.Sprint(userID)).
		AddParam("server", server.String()).
		AddParam("schedule_type", fmt.Sprint(scheduleType)).
		AddDynamicSecret(constants.DS_GLOBAL).
		AddLanguage(genshin.Language).
		Build()

	err = genshin.handler.Send(request, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Get Genshin regional server code based off the user ID.
func (genshin GenshinClient) getRegion(userID int) (*constants.GenshinRegion, error) {
	// Get 1st digit of user ID to determine the region.
	i := userID
	for i >= 10 {
		i /= 10
	}

	region, exists := constants.GenshinServerRegions[i]
	if !exists {
		return nil,
			errors.NewError(
				errors.REGION_SERVER_CODE_ERROR,
				fmt.Sprintf("Genshin regional server code not found for UID %d", userID),
			)
	}

	return &region, nil
}
