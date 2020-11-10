package qoin

var isProduction bool

// SetEnvironment set the API environment
func SetEnvironment(environment string) {
	switch environment {
	case "production":
		isProduction = true
	default:
		isProduction = false
	}
}
