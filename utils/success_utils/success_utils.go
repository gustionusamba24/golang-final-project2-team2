package success_utils

type MessageSuccess interface {
	Message() string
}

type MessageSuccessData struct {
	SuccessMessage string `json:"message"`
}

func (e *MessageSuccessData) Message() string {
	return e.SuccessMessage
}

func Success(message string) MessageSuccess {
	return &MessageSuccessData{
		SuccessMessage: message,
	}
}
