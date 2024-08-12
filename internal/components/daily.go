package components

import (
	"fmt"
	"net/http"

	"github.com/sparkeexd/hoyoapi/handler"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/models"
)

// Daily reward component.
// Responsible for getting information of current daily reward status and claiming daily reward.
type DailyReward struct {
	params   DailyRewardParams
	language string
	handler  handler.Handler
}

// Daily reward endpoints are shared across different games with only minor differences to the URL.
// This struct consolidates the common differences between each game.
type DailyRewardParams struct {
	BaseUrl string
	EventId string
	ActId   string
}

// Constructor.
func NewDailyReward(params DailyRewardParams, language string, handler handler.Handler) DailyReward {
	return DailyReward{
		params:   params,
		language: language,
		handler:  handler,
	}
}

// Constructor.
func NewDailyRewardParams(baseUrl string, eventId string, actId string) DailyRewardParams {
	return DailyRewardParams{
		BaseUrl: baseUrl,
		EventId: eventId,
		ActId:   actId,
	}
}

// Get the list of available daily rewards for the month.
func (daily DailyReward) List() (models.DailyRewardList, error) {
	var res models.DailyRewardList

	endpoint := daily.dailyRewardAPI(daily.params, constants.DAILY_REWARD_HOME)
	request := handler.NewRequest(endpoint, http.MethodGet).
		AddParam("lang", daily.language).
		Build()

	err := daily.handler.Send(request, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Get information on the current daily reward status.
func (daily DailyReward) Info() (models.DailyRewardInfo, error) {
	var res models.DailyRewardInfo

	endpoint := daily.dailyRewardAPI(daily.params, constants.DAILY_REWARD_INFO)
	request := handler.NewRequest(endpoint, http.MethodGet).
		AddParam("lang", daily.language).
		Build()

	err := daily.handler.Send(request, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Claim daily reward.
func (daily DailyReward) Claim() (models.DailyRewardClaim, error) {
	var res models.DailyRewardClaim

	endpoint := daily.dailyRewardAPI(daily.params, constants.DAILY_REWARD_SIGN)
	request := handler.NewRequest(endpoint, http.MethodPost).
		AddParam("lang", daily.language).
		Build()

	err := daily.handler.Send(request, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Returns the endpoint for daily rewards.
func (daily DailyReward) dailyRewardAPI(params DailyRewardParams, endpoint constants.DailyRewardEndpoint) string {
	return fmt.Sprintf("%s/event/%s/%s?act_id=%s", params.BaseUrl, params.EventId, endpoint, params.ActId)
}
