package repository

// Repository - репозиторий
type Repository interface {
	Register() error
	Login() error
}