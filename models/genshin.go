package models

// Genshin Spiral Abyss response structure.
type SpiralAbyss struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		ScheduleID       int    `json:"schedule_id"`
		StartTime        string `json:"start_time"`
		EndTime          string `json:"end_time"`
		TotalBattleTimes int    `json:"total_battle_times"`
		TotalWinTimes    int    `json:"total_win_times"`
		MaxFloor         string `json:"max_floor"`
		RevealRank       []struct {
			AvatarID   int    `json:"avatar_id"`
			AvatarIcon string `json:"avatar_icon"`
			Value      int    `json:"value"`
			Rarity     int    `json:"rarity"`
		} `json:"reveal_rank"`
		DefeatRank []struct {
			AvatarID   int    `json:"avatar_id"`
			AvatarIcon string `json:"avatar_icon"`
			Value      int    `json:"value"`
			Rarity     int    `json:"rarity"`
		} `json:"defeat_rank"`
		DamageRank []struct {
			AvatarID   int    `json:"avatar_id"`
			AvatarIcon string `json:"avatar_icon"`
			Value      int    `json:"value"`
			Rarity     int    `json:"rarity"`
		} `json:"damage_rank"`
		TakeDamageRank []struct {
			AvatarID   int    `json:"avatar_id"`
			AvatarIcon string `json:"avatar_icon"`
			Value      int    `json:"value"`
			Rarity     int    `json:"rarity"`
		} `json:"take_damage_rank"`
		NormalSkillRank []struct {
			AvatarID   int    `json:"avatar_id"`
			AvatarIcon string `json:"avatar_icon"`
			Value      int    `json:"value"`
			Rarity     int    `json:"rarity"`
		} `json:"normal_skill_rank"`
		EnergySkillRank []struct {
			AvatarID   int    `json:"avatar_id"`
			AvatarIcon string `json:"avatar_icon"`
			Value      int    `json:"value"`
			Rarity     int    `json:"rarity"`
		} `json:"energy_skill_rank"`
		Floors []struct {
			Index      int    `json:"index"`
			Icon       string `json:"icon"`
			IsUnlock   bool   `json:"is_unlock"`
			SettleTime string `json:"settle_time"`
			Star       int    `json:"star"`
			MaxStar    int    `json:"max_star"`
			Levels     []struct {
				Index   int `json:"index"`
				Star    int `json:"star"`
				MaxStar int `json:"max_star"`
				Battles []struct {
					Index     int    `json:"index"`
					Timestamp string `json:"timestamp"`
					Avatars   []struct {
						ID     int    `json:"id"`
						Icon   string `json:"icon"`
						Level  int    `json:"level"`
						Rarity int    `json:"rarity"`
					} `json:"avatars"`
					SettleDateTime struct {
						Year   int `json:"year"`
						Month  int `json:"month"`
						Day    int `json:"day"`
						Hour   int `json:"hour"`
						Minute int `json:"minute"`
						Second int `json:"second"`
					} `json:"settle_date_time"`
				} `json:"battles"`
				TopHalfFloorMonster []struct {
					Name  string `json:"name"`
					Icon  string `json:"icon"`
					Level int    `json:"level"`
				} `json:"top_half_floor_monster"`
				BottomHalfFloorMonster []struct {
					Name  string `json:"name"`
					Icon  string `json:"icon"`
					Level int    `json:"level"`
				} `json:"bottom_half_floor_monster"`
			} `json:"levels"`
			SettleDateTime *struct {
				Year   int `json:"year"`
				Month  int `json:"month"`
				Day    int `json:"day"`
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
				Second int `json:"second"`
			} `json:"settle_date_time"`
			LeyLineDisorder []string `json:"ley_line_disorder"`
		} `json:"floors"`
		TotalStar int  `json:"total_star"`
		IsUnlock  bool `json:"is_unlock"`
	} `json:"data"`
}

// HoYoWiki Genshin characters list response structure.
type GenshinCharacters struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			EntryPageID  string `json:"entry_page_id"`
			Name         string `json:"name"`
			IconURL      string `json:"icon_url"`
			DisplayField struct {
			} `json:"display_field"`
			FilterValues struct {
				CharacterRegion struct {
					Values     []string `json:"values"`
					ValueTypes []struct {
						ID         string `json:"id"`
						Value      string `json:"value"`
						Mi18NKey   string `json:"mi18n_key"`
						Icon       string `json:"icon"`
						EnumString string `json:"enum_string"`
					} `json:"value_types"`
					Key any `json:"key"`
				} `json:"character_region,omitempty"`
				CharacterVision struct {
					Values     []string `json:"values"`
					ValueTypes []struct {
						ID         string `json:"id"`
						Value      string `json:"value"`
						Mi18NKey   string `json:"mi18n_key"`
						Icon       string `json:"icon"`
						EnumString string `json:"enum_string"`
					} `json:"value_types"`
					Key any `json:"key"`
				} `json:"character_vision,omitempty"`
				CharacterWeapon struct {
					Values     []string `json:"values"`
					ValueTypes []struct {
						ID         string `json:"id"`
						Value      string `json:"value"`
						Mi18NKey   string `json:"mi18n_key"`
						Icon       string `json:"icon"`
						EnumString string `json:"enum_string"`
					} `json:"value_types"`
					Key any `json:"key"`
				} `json:"character_weapon,omitempty"`
				CharacterProperty struct {
					Values     []string `json:"values"`
					ValueTypes []struct {
						ID         string `json:"id"`
						Value      string `json:"value"`
						Mi18NKey   string `json:"mi18n_key"`
						Icon       string `json:"icon"`
						EnumString string `json:"enum_string"`
					} `json:"value_types"`
					Key any `json:"key"`
				} `json:"character_property,omitempty"`
				CharacterRarity struct {
					Values     []string `json:"values"`
					ValueTypes []struct {
						ID         string `json:"id"`
						Value      string `json:"value"`
						Mi18NKey   string `json:"mi18n_key"`
						Icon       string `json:"icon"`
						EnumString string `json:"enum_string"`
					} `json:"value_types"`
					Key any `json:"key"`
				} `json:"character_rarity,omitempty"`
			} `json:"filter_values"`
			Desc string `json:"desc"`
		}
		Total string `json:"total"`
	} `json:"data"`
}
