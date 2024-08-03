package main

import (
	"net/http"

	"github.com/sparkeexd/hoyoapi/client"
	"github.com/sparkeexd/hoyoapi/handler"
	"github.com/sparkeexd/hoyoapi/internal/constants"
	"github.com/sparkeexd/hoyoapi/internal/utilities"
	"github.com/sparkeexd/hoyoapi/middleware"
)

// Claim Genshin daily rewards.
func GenshinDailyReward(options client.ClientOptions) {
	genshin := client.NewGenshinClient(options)
	response, err := genshin.Daily.Claim()
	utilities.PrintJSON(response, err)
}

// Get Spiral Abyss information.
func GetSpiralAbyssInfo(options client.ClientOptions) {
	genshin := client.NewGenshinClient(options)
	response, err := genshin.SpiralAbyss(true)
	utilities.PrintJSON(response, err)
}

// Get Genshin characters list in HoYoWiki.
func GetGenshinCharacters(cookie middleware.Cookie) {
	request := handler.NewRequest(constants.HOYOWIKI_ENTRY_PAGE_LIST_API, http.MethodPost).
		AddReferer("https://wiki.hoyolab.com").
		AddBody("filters", []string{}).
		AddBody("menu_id", 2).    // Genshin Character List
		AddBody("page_num", 1).   // Pagination
		AddBody("page_size", 30). // Number of items returned
		AddBody("use_es", true).
		Build()

	handler := handler.NewHandler(cookie)

	response, err := handler.Send(request)
	utilities.PrintJSON(response, err)
}
