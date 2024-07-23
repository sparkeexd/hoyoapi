package client

// Client options that are required by each game client i.e., cookies, language and in-game UID.
type ClientOptions struct {
	ltokenV2 string
	ltmidV2  string
	ltuidV2  string
	language string
	userId   int
}

// Client options builder that builds the necessary parameters to use a game client using the builder pattern.
type ClientOptionsBuilder struct {
	options ClientOptions
}

// Constructor.
func NewClientOptions() *ClientOptionsBuilder {
	return &ClientOptionsBuilder{}
}

// Return built client options instance.
func (builder *ClientOptionsBuilder) Build() ClientOptions {
	return ClientOptions{
		ltokenV2: builder.options.ltokenV2,
		ltmidV2:  builder.options.ltmidV2,
		ltuidV2:  builder.options.ltuidV2,
		language: builder.options.language,
		userId:   builder.options.userId,
	}
}

// Add cookie.
func (builder *ClientOptionsBuilder) AddCookie(ltokenV2 string, ltmidV2 string, ltuidV2 string) *ClientOptionsBuilder {
	builder.options.ltokenV2 = ltokenV2
	builder.options.ltmidV2 = ltmidV2
	builder.options.ltuidV2 = ltuidV2
	return builder
}

// Add language.
func (builder *ClientOptionsBuilder) AddLanguage(language string) *ClientOptionsBuilder {
	builder.options.language = language
	return builder
}

// Add user ID.
func (builder *ClientOptionsBuilder) AddUserId(userId int) *ClientOptionsBuilder {
	builder.options.userId = userId
	return builder
}
