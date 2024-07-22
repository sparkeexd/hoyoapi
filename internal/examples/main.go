package main

import (
	"strconv"

	"github.com/sparkeexd/hoyoapi/hoyoapi"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
)

func main() {
	ltokenV2 := utilities.GetEnv("LTOKEN_V2", utilities.ParseString)
	ltmidV2 := utilities.GetEnv("LTMID_V2", utilities.ParseString)
	ltuidV2 := utilities.GetEnv("LTUID_V2", utilities.ParseString)

	language := "en-us"
	userId := utilities.GetEnv("GENSHIN_UID", strconv.Atoi)

	options := hoyoapi.NewClientOptions().
		AddCookie(ltokenV2, ltmidV2, ltuidV2).
		AddLanguage(language).
		AddUserId(userId).
		Build()

	GenshinDailyReward(options)
	StarRailDailyReward(options)
	ZenlessDailyReward(options)
}
