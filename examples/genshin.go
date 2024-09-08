package main

import (
	"strconv"

	"github.com/sparkeexd/hoyoapi/client"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
	"github.com/sparkeexd/hoyoapi/services"
)

// Claim Genshin daily rewards.
func GenshinDailyReward(cookie services.Cookie) {
	genshin := client.NewGenshinClient()
	response, err := genshin.Daily.Claim(cookie)
	utilities.PrintJSON(response, err)
}

// Get Spiral Abyss information.
func GenshinSpiralAbyssInfo(cookie services.Cookie) {
	genshin := client.NewGenshinClient()
	userId := utilities.GetEnv("GENSHIN_UID", strconv.Atoi)

	response, err := genshin.SpiralAbyss(cookie, userId, true)
	utilities.PrintJSON(response, err)
}

// Get Genshin characters list in HoYoWiki.
func GetGenshinCharacters(cookie services.Cookie) {
	genshin := client.NewGenshinClient()
	response, err := genshin.Characters(cookie)
	utilities.PrintJSON(response, err)
}
