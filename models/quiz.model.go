package models

type Quiz struct {
	ID          primitve.ID `json:"id" bson:"_id"`
	UserId      string      `json:"userId" bson:"userId"`
	Title       string      `json:"title" bson:"title"`
	Description string      `json:"description" bson:"description"`
	TimeLimit   string      `json:"timeLimit" bson:"timeLimit"`
	Questions
}
type Questions struct {
	ID            string `json:"id" bson:"id"`
	CorrectAnswer string | []string `json:"correctAnswer" bson:"correctAnswer"`
}
