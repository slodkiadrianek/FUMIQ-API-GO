package schemas

type CreateQuiz struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	TimeLimit   int         `json:"timeLimit" binding:"required"`
	Questions   []Questions `json:"questions" binding:"required"`
}

type Questions struct {
	CorrectAnswer string  `json:"correctAnswer" binding:"required"`
	Options       *string `json:"options"  binding:"required"`
	QuestionText  string  `json:"questionText" binding:"required"`
	QuestionType  string  `json:"questionType" binding:"required"`
}

type QuizParams struct {
	QuizId string `json:"quizId" binding:"required"`
}
