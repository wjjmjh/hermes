package auth_service

type Auth struct {
	Username string
	Password string
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	// todo
	return false, nil
}

func (a *Auth) Check() (bool, error) {
	return CheckAuth(a.Username, a.Password)
}
