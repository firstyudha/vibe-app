package user

// UserService defines the interface for user business logic
type UserService interface {
	Register(request *CreateUserRequest) (*User, error)
	GetByID(id uint) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}
