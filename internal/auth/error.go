package auth

type errorNoAnonymous struct{}

func (e *errorNoAnonymous) ClientError() string {
	return "ErrorAuth"
}

func (e *errorNoAnonymous) Error() string {
	return "Anonymous is not enabled."
}
