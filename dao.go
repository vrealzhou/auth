package auth

// DAO interface for access data storage
type DAO interface {
	GetUserById(id string) (*User, error)
	GetUserByPhone(phone string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
}
