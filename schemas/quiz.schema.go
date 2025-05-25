package schemas

import z "github.com/Oudwins/zog"

type CreateQuiz struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	TimeLimit   int         `json:"timeLimit"`
	Questions   []Questions `json:"questions"`
}

func (c *CreateQuiz) Validate() (z.ZogIssueMap, error) {
	errMap := CreateQuizSchema.Validate(c)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var CreateQuizSchema = z.Struct(z.Schema{
	"title":       z.String().Required(),
	"description": z.String().Required(),
	"timeLimit":   z.Int().Required(),
	"questions":   z.Slice(QuestionSchema).Min(1).Required(),
})

var QuestionSchema = z.Struct(z.Schema{
	"correctAnswer": z.String().Required(),
	"options":       z.String().Optional(),
	"questionText":  z.String().Required(),
	"QuestionType":  z.String().Required(),
})

type Questions struct {
	CorrectAnswer string  `json:"correctAnswer"`
	Options       *string `json:"options"`
	QuestionText  string  `json:"questionText"`
	QuestionType  string  `json:"questionType"`
}

type QuizParams struct {
	QuizId string `json:"quizId"`
}
