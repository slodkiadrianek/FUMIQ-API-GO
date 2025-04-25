package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quiz struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	TimeLimit   string             `json:"timeLimit" bson:"timeLimit"`
	Questions   []Questions        `json:"questions" bson:"questions"`
}
type Questions struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CorrectAnswer *string            `json:"correctAnswer" bson:"correctAnswer"`
	Options       *[]string          `json:"options" bson:"options"`
	QuestionType  string             `json:"questionType" bson:"questionType"`
	QuestionText  string             `json:"questionText" bson:"questionText"`
	PhotoUrl      *string            `json:"photoUrl" bson:"photoUrl"`
}
