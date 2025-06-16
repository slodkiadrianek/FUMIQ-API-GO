package schemas

import z "github.com/Oudwins/zog"

type SessionId struct {
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
}

func (s *SessionId) Validate() (z.ZogIssueMap, error) {
	errMap := JoinQuizSchema.Validate(s)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var SessionIdSchema = z.Struct(z.Schema{
	"sessionId": z.String().Required(),
	"userId":    z.String().Required(),
})

type JoinQuiz struct {
	Code int `json:"code"`
}

func (j *JoinQuiz) Validate() (z.ZogIssueMap, error) {
	errMap := JoinQuizSchema.Validate(j)
	if errMap != nil {
		return errMap, nil
	}
	return nil, nil
}

var JoinQuizSchema = z.Struct(z.Schema{
	"code": z.Int().Required(),
})
