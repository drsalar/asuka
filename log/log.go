package log

import (
	"time"

	log "github.com/sirupsen/logrus"

	"asuka/conf"
)

func Init() {
	switch conf.LogLevel {
	case "d", "debug":
		log.SetLevel(log.DebugLevel)
	case "i", "info":
		log.SetLevel(log.InfoLevel)
	case "w", "warning":
		log.SetLevel(log.WarnLevel)
	case "e", "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func Info(function string, step string, msg string, infos ...string) {
	fields := makeFields(function, step, msg, infos)
	log.WithFields(fields).Info(msg)
}

func Warning(function string, step string, msg string, infos ...string) {
	fields := makeFields(function, step, msg, infos)
	log.WithFields(fields).Warning(msg)
}

func Error(function string, step string, msg string, infos ...string) {
	fields := makeFields(function, step, msg, infos)
	log.WithFields(fields).Error(msg)
}

func Debug(function string, step string, msg string, infos ...string) {
	fields := makeFields(function, step, msg, infos)
	log.WithFields(fields).Debug(msg)
}

func makeFields(function string, step string, msg string, infos []string) (fields log.Fields) {
	Fields = make(map[string]interface{})
	fields["func"] = function
	fields["step"] = step
	fields["time"] = time.Now().Format("2006-01-02 15:04:05")
	for i := 0; i < len(infos)-1; i += 2 {
		fields[infos[i]] = infos[i+1]
	}
	return
}
