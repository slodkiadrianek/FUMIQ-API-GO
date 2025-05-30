package models

import (
	"FUMIQ_API/schemas"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quiz struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	TimeLimit   int                `json:"timeLimit" bson:"timeLimit"`
	Questions   []Question         `json:"questions" bson:"questions"`
}

type Question struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CorrectAnswer string             `json:"correctAnswer" bson:"correctAnswer"`
	Options       []string           `json:"options" bson:"options"`
	QuestionType  string             `json:"questionType" bson:"questionType"`
	QuestionText  string             `json:"questionText" bson:"questionText"`
	PhotoUrl      string             `json:"photoUrl" bson:"photoUrl"`
}

func newQuestion(correctAnswer string, questionText, questionType string, photoUrl string, options []string) Question {
	return Question{
		ID:            primitive.NewObjectID(),
		CorrectAnswer: correctAnswer,
		Options:       options,
		QuestionType:  questionType,
		QuestionText:  questionText,
		PhotoUrl:      photoUrl,
	}
}

func NewQuiz(userId, title, description string, timeLimit int, question []schemas.Question) (Quiz, error) {
	// Convert userId string to ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return Quiz{}, err
	}

	quiz := Quiz{
		ID:          primitive.NewObjectID(),
		UserId:      userObjectID,
		Title:       title,
		Description: description,
		TimeLimit:   timeLimit,
	}

	for _, v := range question {
		quiz.Questions = append(quiz.Questions, newQuestion(v.CorrectAnswer, v.QuestionText, v.QuestionType, v.PhotoUrl, v.Options))
	}

	return quiz, nil
}

