package constants

// Daily reward endpoint parameter.
type DailyRewardEndpoint string

// Daily reward endpoints are shared across different games with only minor differences to the URL.
// This struct consolidates the common differences between each game.
type dailyRewardParams struct {
	BaseUrl string
	EventId string
	ActId   string
}

const (
	// Daily reward endpoint parameters.
	DAILY_REWARD_HOME DailyRewardEndpoint = "home"
	DAILY_REWARD_INFO DailyRewardEndpoint = "info"
	DAILY_REWARD_SIGN DailyRewardEndpoint = "sign"

	// Genshin endpoint parameters.
	genshinEventId = "sol"
	genshinActId   = "e202102251931481"

	// Star Rail endpoint parameters.
	starRailEventId = "luna/os"
	starRailActId   = "e202303301540311"

	// Zenless endpoint parameters.
	zenlessEventId = "luna/zzz/os"
	zenlessActId   = "e202406031448091"
)

// Specific endpoint parameters for each game.
var DailyEndpointParams = map[Game]dailyRewardParams{
	GAME_GENSHIN: {
		BaseUrl: HK4E_API,
		EventId: genshinEventId,
		ActId:   genshinActId,
	},
	GAME_STAR_RAIL: {
		BaseUrl: PUBLIC_API,
		EventId: starRailEventId,
		ActId:   starRailActId,
	},
	GAME_ZENLESS: {
		BaseUrl: ACT_API,
		EventId: zenlessEventId,
		ActId:   zenlessActId,
	},
}

// Stringer.
func (endpoint DailyRewardEndpoint) String() string {
	return string(endpoint)
}
