package constants

// Genshin regional server codes.
type GenshinRegion string

const (
	// Genshin regional server codes.
	GENSHIN_REGION_USA          GenshinRegion = "os_usa"
	GENSHIN_REGION_EUROPE       GenshinRegion = "os_euro"
	GENSHIN_REGION_ASIA         GenshinRegion = "os_asia"
	GENSHIN_REGION_CHINA_TAIWAN GenshinRegion = "os_cht"
)

// Mapping of Genshin's UID to server region
var GenshinServerRegions = map[int]GenshinRegion{
	6: GENSHIN_REGION_USA,
	7: GENSHIN_REGION_EUROPE,
	8: GENSHIN_REGION_ASIA,
	9: GENSHIN_REGION_CHINA_TAIWAN,
}

// Stringer.
func (region GenshinRegion) String() string {
	return string(region)
}
