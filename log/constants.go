package log

// Log keys
const (
	StatusCodeKey = "statusCode"
	LatencyKey    = "latency"
	ClientIPKey   = "clientIP"
	MethodKey     = "method"
	PathKey       = "path"
	ErrorKey      = "error"

	HTTPConfigKey     = "httpConfig"
	DatabaseConfigKey = "databaseConfig"
	LogLevelKey       = "logLevel"
	IDLogParam        = "id"
)

// Log Levels
const (
	TraceLevel        = "trace"
	DebugLevel        = "debug"
	InfoLevel         = "info"
	WarnLevel         = "warn"
	ErrorLevel        = "error"
	FatalLevel        = "fatal"
	PanicLevel        = "panic"
	LoggerStorage     = "storage/logger"
	LoggerStorageFile = "logger"
)
