// Different implementations constants, used in configurations for determining the right injection.
package impl

// Available Data Service implementations
type DataServiceImpl string

func (d *DataServiceImpl) Decode(s string) error {
	*d = DataServiceImpl(s)
	return nil
}

var (
	DataServicePostgres DataServiceImpl = "postgresds"
)

// Available Adapter implementations.
type AdapterImpl string

func (a *AdapterImpl) Decode(s string) error {
	*a = AdapterImpl(s)
	return nil
}

var (
	RestAdapter AdapterImpl = "rest"
)

// Available Log Levels.
type LoggerLevel string

func (l *LoggerLevel) Decode(s string) error {
	*l = LoggerLevel(s)
	return nil
}

var (
	TraceLevel LoggerLevel = "trace"
	DebugLevel LoggerLevel = "debug"
	InfoLevel  LoggerLevel = "info"
	WarnLevel  LoggerLevel = "warn"
	ErrorLevel LoggerLevel = "error"
	FatalLevel LoggerLevel = "fatal"
	PanicLevel LoggerLevel = "panic"
)
