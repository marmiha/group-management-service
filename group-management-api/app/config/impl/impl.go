// Different implementations constants, used in configurations for determining the right injection.
package impl

// Available Data Service implementations
type DataServiceImpl string

var (
	DataServicePostgres DataServiceImpl = "postgresds"
)

// Available Adapter implementations.
type AdapterImpl string

var (
	RestAdapter AdapterImpl = "rest"
)

// Available Log Levels.
type LoggerLevel string

var (
	TraceLevel LoggerLevel = "trace"
	DebugLevel LoggerLevel = "debug"
	InfoLevel  LoggerLevel = "info"
	WarnLevel  LoggerLevel = "warn"
	ErrorLevel LoggerLevel = "error"
	FatalLevel LoggerLevel = "fatal"
	PanicLevel LoggerLevel = "panic"
)
