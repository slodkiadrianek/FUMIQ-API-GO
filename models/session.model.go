package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Session struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userId" bson:"userId"`
	QuizID      primitive.ObjectID `json:"quizId" bson:"quizId"`
	Code        string             `json:"code" bson:"code"`
	IsActive    bool               `json:"isActive" bson:"isActive"`
	Competitors []Competitors      `json:"competitors" bson:"competitors"`
}

type Competitors struct {
	UserID   primitive.ObjectID `json:"userId" bson:"userId"`
	Answers  []Answers          `json:"answers" bson:"answers"`
	Finished bool               `json:"finished" bson:"finished"`
}

type Answers struct {
	QuestionId primitive.ObjectID `json:"questionId" bson:"questionId"`
	Answers    primitive.ObjectID `json:"answers" bson:"answers"`
}
