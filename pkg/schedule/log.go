package schedule

type Logger interface {
	Log(level int, typ int, log string)
	D(typ int, log string)
	E(typ int, log string)
	Err(typ int, err error)
	I(typ int, log string)
}

func GetLogger() *Logger {
	return nil
}
