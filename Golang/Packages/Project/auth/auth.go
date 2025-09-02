package auth

// Public function as function start with Capital letter
func GetSession() string {
	return extractSession()
}

// Private function as function start with small letter
func extractSession() string {
	return "loggenIn"
}
