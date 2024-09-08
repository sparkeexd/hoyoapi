package main

import (
	"github.com/sparkeexd/hoyoapi/internal/utilities"
	"github.com/sparkeexd/hoyoapi/services"
)

func main() {
	ltokenV2 := utilities.GetEnv("LTOKEN_V2", utilities.ParseString)
	ltmidV2 := utilities.GetEnv("LTMID_V2", utilities.ParseString)
	ltuidV2 := utilities.GetEnv("LTUID_V2", utilities.ParseString)

	cookie := services.NewCookie(ltokenV2, ltmidV2, ltuidV2)

	GenshinDailyReward(cookie)
	StarRailDailyReward(cookie)
	ZenlessDailyReward(cookie)
}
