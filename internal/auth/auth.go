package auth

type AuthManager struct {
	UserID uint
	Role   string
}

var authInstance *AuthManager

func GetAuthInstance() *AuthManager {
	if authInstance == nil {
		authInstance = &AuthManager{
			UserID: 2,
			Role:   "модератор",
		}
	}
	return authInstance
}
