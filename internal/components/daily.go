package components

import (
	"fmt"
	"net/http"

	"github.com/sparkeexd/hoyoapi/constants"
	"github.com/sparkeexd/hoyoapi/handler"
)

// Daily reward component.
// Responsible for getting information of current daily reward status and claiming daily reward.
type DailyReward struct {
	game     constants.Game
	language string
	handler  handler.Handler
}

// Constructor.
func NewDailyReward(game constants.Game, language string, handler handler.Handler) DailyReward {
	return DailyReward{
		game:     game,
		language: language,
		handler:  handler,
	}
}

// Get the list of available daily rewards for the month.
func (daily DailyReward) List() (map[string]interface{}, error) {
	endpoint := daily.dailyRewardAPI(daily.game, constants.DAILY_REWARD_HOME)
	request := handler.NewRequest(endpoint, http.MethodGet).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.handler.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Get information on the current daily reward status.
func (daily DailyReward) Info() (map[string]interface{}, error) {
	endpoint := daily.dailyRewardAPI(daily.game, constants.DAILY_REWARD_INFO)
	request := handler.NewRequest(endpoint, http.MethodGet).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.handler.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Claim daily reward.
func (daily DailyReward) Claim() (map[string]interface{}, error) {
	endpoint := daily.dailyRewardAPI(daily.game, constants.DAILY_REWARD_SIGN)
	request := handler.NewRequest(endpoint, http.MethodPost).
		AddParam("lang", daily.language).
		Build()

	data, err := daily.handler.Send(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Returns the endpoint for daily rewards.
func (daily DailyReward) dailyRewardAPI(game constants.Game, endpoint constants.DailyRewardEndpoint) string {
	params := constants.DailyEndpointParams[game]
	return fmt.Sprintf("%s/event/%s/%s?act_id=%s", params.BaseUrl, params.EventId, endpoint, params.ActId)
}
