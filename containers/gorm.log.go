package containers

import (
	`context`
	`errors`
	`fmt`
	`time`
	
	gormlogger `gorm.io/gorm/logger`
	`gorm.io/gorm/utils`
)

const (
	// Silent silent log level
	Silent gormlogger.LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
)

const prefix = "[GORM] "

func NewGormLog(writer gormlogger.Writer, config gormlogger.Config) gormlogger.Interface {
	var (
		infoStr      = "%s[INFO] "
		warnStr      = "%s[WARN] "
		errStr       = "%s[ERROR] "
		traceStr     = "%s [用时: %.3fms] [rows:%v]\n\t[sql] %s"
		traceWarnStr = "%s %s [用时: %.3fms] [rows:%v]\n\t[sql] %s"
		traceErrStr  = "%s %s [用时: %.3fms] [rows:%v]\n\t[sql] %s"
	)
	
	if config.Colorful {
		infoStr = gormlogger.Green + "%s" + gormlogger.Reset + gormlogger.Green + " [INFO] " + gormlogger.Reset
		warnStr = gormlogger.BlueBold + "%s" + gormlogger.Reset + gormlogger.Magenta + " [WARN] " + gormlogger.Reset
		errStr = gormlogger.Magenta + "%s" + gormlogger.Reset + gormlogger.Red + " [ERROR] " + gormlogger.Reset
		traceStr = gormlogger.Green + "%s" + gormlogger.Reset + gormlogger.Yellow + " [用时: %.3fms] " + gormlogger.BlueBold + " [rows:%v]" + "\t[sql] " + gormlogger.MagentaBold + " %s" + gormlogger.Reset
		traceWarnStr = gormlogger.Green + "%s " + gormlogger.Yellow + "%s" + gormlogger.Reset + gormlogger.RedBold + " [用时: %.3fms] " + gormlogger.Yellow + " [rows:%v]" + gormlogger.Magenta + " \n\t[sql] %s" + gormlogger.Reset
		traceErrStr = gormlogger.RedBold + "%s " + gormlogger.MagentaBold + "%s" + gormlogger.Reset + gormlogger.Yellow + "[用时: %.3fms] " + gormlogger.BlueBold + " [rows:%v]" + gormlogger.Reset + "\n\t[sql] %s"
	}
	
	return &mgromlogger{
		Writer:       writer,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type mgromlogger struct {
	gormlogger.Writer
	gormlogger.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (l *mgromlogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l mgromlogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Info {
		l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l mgromlogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l mgromlogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= Error {
		l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l mgromlogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= Silent {
		return
	}
	
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= Error && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceErrStr, "", err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceErrStr, "", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Printf(l.traceWarnStr, "", slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceWarnStr, "", slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == Info:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceStr, "", float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceStr, "", float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
