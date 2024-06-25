package model

type ErrorMessage struct {
	Status  int    `bson:"status,omitempty" json:"status"`
	Message string `bson:"message,omitempty" json:"message"`
}