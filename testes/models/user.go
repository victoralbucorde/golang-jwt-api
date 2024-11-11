package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserType     *string            `json:"userType"`
	UserID       string             `json:"userID"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refreshToken"`
	FirstName    *string            `json:"firstName"`
	LastName     *string            `json:"lastName"`
	Email        *string            `json:"email" validate:"required,email"`
	Password     *string            `json:"password" validate:"required,min=8"`
	Save         *Save              `json:"save"`
}
