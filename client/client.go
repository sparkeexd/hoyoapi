package client

// Client options that are required by each game client i.e., cookies, language and in-game UID.
type ClientOptions struct {
	language string
	userID   int
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
		language: builder.options.language,
		userID:   builder.options.userID,
	}
}

// Add language.
func (builder *ClientOptionsBuilder) AddLanguage(language string) *ClientOptionsBuilder {
	builder.options.language = language
	return builder
}

// Add user ID.
func (builder *ClientOptionsBuilder) AddUserID(userID int) *ClientOptionsBuilder {
	builder.options.userID = userID
	return builder
}
