package models

type LoginRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type RegisterRequest struct {
	Email    string `bson:"email"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
}
