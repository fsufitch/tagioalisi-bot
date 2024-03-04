package log

import "github.com/google/wire"

// CLILoggingProviderSet provides everything needed to log to the CLI
var CLILoggingProviderSet = wire.NewSet(
	ProvideLogger,
	ProvideStdOutReceiver,
	ProvideStdErrReceiver,
	wire.Struct(new(CLILoggingBootstrapper), "*"),
)
