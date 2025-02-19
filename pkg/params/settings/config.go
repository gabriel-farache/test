package settings

const (
	CliBinaryName         = "go-kcloutie"
	DebugModeLoggerEnvVar = "GOKCLOUTIE_DEBUG"
)

var (
	RootOptions      = RootFlags{}
	DebugModeEnabled = false
	IsQuiet          = false
)

type RootFlags struct {
	NoColor bool
}
