package models

type User struct {
    ID       string `json:"id"`
    Email    string `json:"email"`
    Phone    string `json:"phone"`
    Password string `json:"password"`
}
