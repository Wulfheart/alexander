package logging

import "go.uber.org/zap"

var(
	Logger *zap.Logger
)

func Instantiate(l *zap.Logger) {
	Logger = l
}