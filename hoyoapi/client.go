package hoyoapi

// Client options that are required by each game client i.e., cookies, language and in-game UID.
type ClientOptions struct {
	ltokenV2 string
	ltmidV2  string
	ltuidV2  string
	language string
	userId   int
}

// Create a struct of client options to be passed to a game client.
func NewClientOptions(ltokenV2 string, ltmidV2 string, ltuidV2 string, language string, userId int) ClientOptions {
	return ClientOptions{
		ltokenV2: ltokenV2,
		ltmidV2:  ltmidV2,
		ltuidV2:  ltuidV2,
		language: language,
		userId:   userId,
	}
}
