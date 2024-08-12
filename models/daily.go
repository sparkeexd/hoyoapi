package models

/// Daily reward list response structure.
type DailyRewardList struct {
	Data struct {
		Awards []struct {
			Cnt  int    `json:"cnt"`
			Icon string `json:"icon"`
			Name string `json:"name"`
		} `json:"awards"`
		Month  int    `json:"month"`
		Now    string `json:"now"`
		Resign bool   `json:"resign"`
	} `json:"data"`
	Message string `json:"message"`
	Retcode int    `json:"retcode"`
}

/// Daily reward info response structure.
type DailyRewardInfo struct {
	Data struct {
		FirstBind    bool   `json:"first_bind"`
		IsSign       bool   `json:"is_sign"`
		IsSub        bool   `json:"is_sub"`
		MonthLastDay bool   `json:"month_last_day"`
		Region       string `json:"region"`
		Today        string `json:"today"`
		TotalSignDay int    `json:"total_sign_day"`
	} `json:"data"`
	Message string `json:"message"`
	Retcode int    `json:"retcode"`
}

/// Daily reward claim response structure.
type DailyRewardClaim struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
}
