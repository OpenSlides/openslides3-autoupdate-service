package auth

// Error is returned, when something goes wrong in the auth code.
type Error string

func (e Error) Error() string {
	if e == "" {
		return "Authentication failed"
	}
	return string(e)
}
