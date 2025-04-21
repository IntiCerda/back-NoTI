// internal/models/user.go
package models

// User representa un usuario en nuestro sistema
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
