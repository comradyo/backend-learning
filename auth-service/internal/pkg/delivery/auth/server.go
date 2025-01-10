package auth

// Server - интерфейс обработки запросов регистрации/авторизации
type Server interface {
	Register() error
}
