package schemas

import z "github.com/Oudwins/zog"

type SessionParams struct {
	SessionId string `json:"sessionId" binding:"required"`
}

type JoinQuiz struct {
	Code string `json:"code"`
}

func (j *JoinQuiz) Validate() (z.ZogIssueMap, error) {
	errMap := JoinQuizSchema.Validate(j)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var JoinQuizSchema = z.Struct(z.Schema{
	"code": z.String().Required(),
})
