package constants

// Daily reward endpoint parameter.
type DailyRewardEndpoint string

const (
	// Daily reward endpoint parameters.
	DAILY_REWARD_HOME DailyRewardEndpoint = "home"
	DAILY_REWARD_INFO DailyRewardEndpoint = "info"
	DAILY_REWARD_SIGN DailyRewardEndpoint = "sign"

	// Genshin endpoint parameters.
	GenshinEventId = "sol"
	GenshinActId   = "e202102251931481"

	// Star Rail endpoint parameters.
	StarRailEventId = "luna/os"
	StarRailActId   = "e202303301540311"

	// Zenless endpoint parameters.
	ZenlessEventId = "luna/zzz/os"
	ZenlessActId   = "e202406031448091"
)

// Stringer.
func (endpoint DailyRewardEndpoint) String() string {
	return string(endpoint)
}
