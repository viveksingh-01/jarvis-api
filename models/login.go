package models

type LoginRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
