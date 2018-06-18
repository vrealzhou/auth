package auth

// Service for business logic in user
type Service struct {
}

// Login is for user login
func (s *Service) Login(loginType LoginType, username, password string) (*User, error) {
	return nil, nil
}

// GetUser get user by userID
func (s *Service) GetUser(id string) (*User, error) {
	return nil, nil
}

// UpdateUser is for update user
func (s *Service) UpdateUser(user *User) (*User, error) {
	return nil, nil
}

// DeleteUser is for delete user
func (s *Service) DeleteUser(id string) (*User, error) {
	return nil, nil
}
