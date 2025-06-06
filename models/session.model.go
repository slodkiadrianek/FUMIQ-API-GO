package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userId" bson:"userId"`
	QuizID      primitive.ObjectID `json:"quizId" bson:"quizId"`
	Code        int                `json:"code" bson:"code"`
	IsActive    bool               `json:"isActive" bson:"isActive"`
	Competitors []Competitors      `json:"competitors" bson:"competitors"`
	StartedAt   time.Time          `json:"startedAt" bson:"startedAt"`
	EndedAt     time.Time          `json:"endedAt" bson:"endedAt"`
}

type Competitors struct {
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	Answers   []Answers          `json:"answers" bson:"answers"`
	StartedAt *time.Time         `json:"startedAt" bson:"startedAt"`
	Finished  bool               `json:"finished" bson:"finished"`
}

type Answers struct {
	QuestionId primitive.ObjectID `json:"questionId" bson:"questionId"`
	Answer     string             `json:"answer" bson:"answer"`
}

func NewSession(userId, quizId primitive.ObjectID, code int) Session {
	return Session{
		ID:          primitive.NewObjectID(),
		UserID:      userId,
		QuizID:      quizId,
		Code:        code,
		IsActive:    true,
		Competitors: []Competitors{},
		StartedAt:   time.Now(),
		EndedAt:     time.Now(),
	}
}
