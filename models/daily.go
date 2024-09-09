package models

type DailyRewardList struct {
	Data    DailyRewardListData `json:"data"`
	Message string              `json:"message"`
	Retcode int                 `json:"retcode"`
}

type DailyRewardListData struct {
	Awards DailyRewardAward `json:"awards"`
	Month  int              `json:"month"`
	Now    string           `json:"now"`
	Resign bool             `json:"resign"`
}

type DailyRewardAward struct {
	Cnt  int    `json:"cnt"`
	Icon string `json:"icon"`
	Name string `json:"name"`
}

type DailyRewardInfo struct {
	Data    DailyRewardInfoData `json:"data"`
	Message string              `json:"message"`
	Retcode int                 `json:"retcode"`
}

type DailyRewardInfoData struct {
	FirstBind    bool   `json:"first_bind"`
	IsSign       bool   `json:"is_sign"`
	IsSub        bool   `json:"is_sub"`
	MonthLastDay bool   `json:"month_last_day"`
	Region       string `json:"region"`
	Today        string `json:"today"`
	TotalSignDay int    `json:"total_sign_day"`
}

type DailyRewardClaim struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
}
