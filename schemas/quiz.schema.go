package schemas

import z "github.com/Oudwins/zog"

type CreateQuiz struct {
	Title       string     `json:"title"`
	UserId      string     `json:"userId"`
	Description string     `json:"description"`
	TimeLimit   int        `json:"timeLimit"`
	Questions   []Question `json:"questions"`
}

func (c *CreateQuiz) Validate() (z.ZogIssueMap, error) {
	errMap := CreateQuizSchema.Validate(c)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var CreateQuizSchema = z.Struct(z.Shape{
	"title":       z.String().Required(),
	"userId":      z.String().Required(),
	"description": z.String().Required(),
	"timeLimit":   z.Int().Required(),
	"questions":   z.Slice(QuestionSchema).Min(1).Required(),
})

var QuestionSchema = z.Struct(z.Shape{
	"correctAnswer": z.String().Required(),
	"options":       z.Slice(z.String()).Optional(),
	"questionText":  z.String().Required(),
	"questionType":  z.String().Required(),
	"photoUrl":      z.String().Required(),
})

type Question struct {
	CorrectAnswer string    `json:"correctAnswer"`
	Options       *[]string `json:"options"`
	QuestionText  string    `json:"questionText"`
	QuestionType  string    `json:"questionType"`
	PhotoUrl      string    `json:"photoUrl"`
}

type QuizParams struct {
	QuizId string `json:"quizId"`
}

func (q *QuizParams) Validate() (z.ZogIssueMap, error) {
	errMap := QuizParamsSchema.Validate(q)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var QuizParamsSchema = z.Struct(z.Shape{
	"quizId": z.String().Required(),
})

type EndSession struct {
	SessionId string `json:"sessionId"`
	QuizId    string `json:"quizId"`
}

func (e *EndSession) Validate() (z.ZogIssueMap, error) {
	errMap := EndSessionSchema.Validate(e)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var EndSessionSchema = z.Struct(z.Shape{
	"sessionId": z.String().Required(),
	"quizId":    z.String().Required(),
})
