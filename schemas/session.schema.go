package schemas

type SessionParams struct {
	SessionId string `json:"sessionId" binding:"required"`
}
