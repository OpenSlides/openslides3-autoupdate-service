package auth

// NoAnonymousError is returned when a anonymous tries to connect but anonymous is
// disabled.
type NoAnonymousError struct{}

func (e NoAnonymousError) Error() string {
	return "Anonymous is not enabled."
}
