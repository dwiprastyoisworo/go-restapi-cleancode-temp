package repositories

type UserRepository struct {
}

func NewUserRepository() UserRepositoryImpl {
	return &UserRepository{}
}

type UserRepositoryImpl interface {
}
