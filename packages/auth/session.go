package auth

// By using a lowercase first letter, this function is unexported (private) to the package
func extractSessionToken() string {
	return "SESSION_TOKEN_ABC123"
}

func GetSessionToken() string {
	return extractSessionToken()
}
