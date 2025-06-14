package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionQuestions struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Quiz       Quiz               `json:"quiz" bson:"quiz"`
	Competitor Competitor         `json:"competitor" bson:"competitor"`
}

type QuestionSession struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	UserID     primitive.ObjectID `json:"userId" `
	QuizID     primitive.ObjectID `json:"quizId"`
	Competitor Competitor         `json:"competitor"`
}

type PopulatedSession struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userId" bson:"userId"`
	QuizID      Quiz               `json:"quizId" bson:"quizId"`
	Code        int                `json:"code" bson:"code"`
	IsActive    bool               `json:"isActive" bson:"isActive"`
	Competitors []Competitor       `json:"competitors" bson:"competitors"`
	StartedAt   time.Time          `json:"startedAt" bson:"startedAt"`
	EndedAt     time.Time          `json:"endedAt" bson:"endedAt"`
}

type PopulatedCompetitors struct {
	UserID    User       `json:"userId" bson:"userId"`
	Answers   []Answers  `json:"answers" bson:"answers"`
	StartedAt *time.Time `json:"startedAt" bson:"startedAt"`
	Finished  bool       `json:"finished" bson:"finished"`
}

type SessionInfo struct {
	StartedAt            time.Time `json:"startedAt"`
	EndedAt              time.Time `json:"endedAt"`
	AmountOfParticipants int       `json:"AmountOfParticipants"`
}

type Results struct {
	Name        string       `json:"name"`
	Score       int          `json:"score"`
	UserAnswers []UserAnswer `json:"userAnswers"`
}

type UserAnswer struct {
	QuestionText string `json:"questionText"`
	Answer       string `json:"asnwer"`
}

type Session struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"userId" bson:"userId"`
	QuizID      primitive.ObjectID `json:"quizId" bson:"quizId"`
	Code        int                `json:"code" bson:"code"`
	IsActive    bool               `json:"isActive" bson:"isActive"`
	Competitors []Competitor       `json:"competitors" bson:"competitors"`
	StartedAt   time.Time          `json:"startedAt" bson:"startedAt"`
	EndedAt     time.Time          `json:"endedAt" bson:"endedAt"`
}

type Competitor struct {
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	Answers   []Answers          `json:"answers" bson:"answers"`
	StartedAt time.Time          `json:"startedAt" bson:"startedAt"`
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
		Competitors: []Competitor{},
		StartedAt:   time.Now(),
		EndedAt:     time.Now(),
	}
}
