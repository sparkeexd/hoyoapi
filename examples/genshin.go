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
	userID := utilities.GetEnv("GENSHIN_UID", strconv.Atoi)

	genshin := client.NewGenshinClient()
	response, err := genshin.SpiralAbyss(cookie, userID, true)
	utilities.PrintJSON(response, err)
}
