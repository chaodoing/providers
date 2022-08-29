package containers

import (
	`strings`
	
	`gorm.io/gorm/logger`
)

func gormLevel(level string) logger.LogLevel {
	var data = map[string]logger.LogLevel{
		"silent": logger.Silent,
		"error":  logger.Error,
		"warn":   logger.Warn,
		"info":   logger.Info,
	}
	if val, ok := data[strings.ToLower(level)]; ok {
		return val
	}
	return logger.Warn
}

func irisLevel(level string) string {
	var data = map[string]string{
		"disable": "disable",
		"fatal":   "fatal",
		"error":   "error",
		"warn":    "warn",
		"info":    "info",
		"debug":   "debug",
	}
	if val, ok := data[strings.ToLower(level)]; ok {
		return val
	}
	return "disable"
}
