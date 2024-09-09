package models

type SpiralAbyss struct {
	Data    Nullable[SpiralAbyssData] `json:"data"`
	Message string                    `json:"message"`
	Retcode int                       `json:"retcode"`
}

type SpiralAbyssData struct {
	ScheduleID       int                `json:"schedule_id"`
	StartTime        string             `json:"start_time"`
	EndTime          string             `json:"end_time"`
	TotalBattleTimes int                `json:"total_battle_times"`
	TotalWinTimes    int                `json:"total_win_times"`
	MaxFloor         string             `json:"max_floor"`
	RevealRank       []SpiralAbyssRank  `json:"reveal_rank"`
	DefeatRank       []SpiralAbyssRank  `json:"defeat_rank"`
	DamageRank       []SpiralAbyssRank  `json:"damage_rank"`
	TakeDamageRank   []SpiralAbyssRank  `json:"take_damage_rank"`
	NormalSkillRank  []SpiralAbyssRank  `json:"normal_skill_rank"`
	EnergySkillRank  []SpiralAbyssRank  `json:"energy_skill_rank"`
	Floors           []SpiralAbyssFloor `json:"floors"`
	TotalStar        int                `json:"total_star"`
	IsUnlock         bool               `json:"is_unlock"`
}

type SpiralAbyssRank struct {
	AvatarID   int    `json:"avatar_id"`
	AvatarIcon string `json:"avatar_icon"`
	Value      int    `json:"value"`
	Rarity     int    `json:"rarity"`
}

type SpiralAbyssFloor struct {
	Index           int                                 `json:"index"`
	Icon            string                              `json:"icon"`
	IsUnlock        bool                                `json:"is_unlock"`
	SettleTime      string                              `json:"settle_time"`
	Star            int                                 `json:"star"`
	MaxStar         int                                 `json:"max_star"`
	Levels          []SpiralAbyssLevel                  `json:"levels"`
	SettleDateTime  Nullable[SpiralAbyssSettleDateTime] `json:"settle_date_time"`
	LeyLineDisorder []string                            `json:"ley_line_disorder"`
}

type SpiralAbyssLevel struct {
	Index                  int                       `json:"index"`
	Star                   int                       `json:"star"`
	MaxStar                int                       `json:"max_star"`
	Battles                []SpiralAbyssBattle       `json:"battles"`
	TopHalfFloorMonster    []SpiralAbyssFloorMonster `json:"top_half_floor_monster"`
	BottomHalfFloorMonster []SpiralAbyssFloorMonster `json:"bottom_half_floor_monster"`
}

type SpiralAbyssBattle struct {
	Index          int                       `json:"index"`
	Timestamp      string                    `json:"timestamp"`
	Avatars        []SpiralAbyssAvatar       `json:"avatars"`
	SettleDateTime SpiralAbyssSettleDateTime `json:"settle_date_time"`
}

type SpiralAbyssAvatar struct {
	ID     int    `json:"id"`
	Icon   string `json:"icon"`
	Level  int    `json:"level"`
	Rarity int    `json:"rarity"`
}

type SpiralAbyssSettleDateTime struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

type SpiralAbyssFloorMonster struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Level int    `json:"level"`
}
