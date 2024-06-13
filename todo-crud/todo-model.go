package todocrud

type ToDoInfo struct {
	Title       string `bson:"title,omitempty" json:"title"`
	Description string `bson:"description,omitempty" json:"description"`
	CreatedAt   string `bson:"created_at,omitempty" json:"created_at"`
}

type ErrorMessage struct {
	Message string `bson:"message,omitempty" json:"message"`
}
