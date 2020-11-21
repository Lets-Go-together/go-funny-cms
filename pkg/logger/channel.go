package logger

// Logger channel
// --------------

func Info(title string, content string) {
	handle("Info", title, content)
}

func Error(title string, content string) {
	handle("Error", title, content)
}
