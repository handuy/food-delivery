package common

type AppError struct {
	StatusCode int       `json:"status_code"`
	RootErr    error     `json:"-"`
	Message    string    `json:"message"`
	Log        string    `json:"log"`
	Key        string    `json:"key"`
}

func NewErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func(a *AppError) Error() string {
	return a.RootErr.Error()
}