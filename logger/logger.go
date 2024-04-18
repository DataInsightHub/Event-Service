package logger

type (
	Logger interface {
		Info(string)
		Warning(string)
		Error(string)
	}
)