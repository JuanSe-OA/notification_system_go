// dto/login_request.go
package dto

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
