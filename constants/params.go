package constants

// Game types in HoYoLab.
type Game string

// Language codes supported in HoYoLab.
type Language string

// Dynamic secret salts used in HoYoLab.
type DynamicSecret string

const (
	// Game types.
	GAME_GENSHIN   Game = "GENSHIN"
	GAME_STAR_RAIL Game = "STAR_RAIL"
	GAME_ZENLESS   Game = "ZENLESS"

	// Languages.
	LANG_SIMPLIFIED_CHINESE Language = "zh-cn"
	LANG_TRADIIONAL_CHINESE Language = "zh-tw"
	LANG_GERMAN             Language = "de-de"
	LANG_ENGLISH            Language = "en-us"
	LANG_SPANISH            Language = "es-es"
	LANG_FRENCH             Language = "fr-fr"
	LANG_INDONESIAN         Language = "id-id"
	LANG_ITALIAN            Language = "it-it"
	LANG_JAPANESE           Language = "ja-jp"
	LANG_KOREAN             Language = "ko-kr"
	LANG_PORTUGUESE         Language = "pt-pt"
	LANG_RUSSIAN            Language = "ru-ru"
	LANG_THAI               Language = "th-th"
	LANG_TURKISH            Language = "tr-tr"
	LANG_VIETNAMESE         Language = "vi-vn"

	// Dynamic secret salts.
	DS_GLOBAL    DynamicSecret = "6s25p5ox5y14umn1p61aqyyvbvvl3lrt"
	DS_APP_LOGIN DynamicSecret = "IZPgfb0dRPtBeLuFkdDznSZ6f4wWt6y2"
)

// Stringer.
func (game Game) String() string {
	return string(game)
}

// Stringer.
func (language Language) String() string {
	return string(language)
}

// Stringer.
func (secret DynamicSecret) String() string {
	return string(secret)
}
