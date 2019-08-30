package cock

import (
	"io"
	"os"
)

// EnvCockMode indicates environment name for cock mode.
const EnvCockMode = "COCK_MODE"

const (
	// DebugMode indicates cock mode is debug.
	DebugMode = "debug"
	// ReleadeMode indicates cock mode is release.
	ReleaseMode = "release"
	// TestMode indicates cock mode is test.
	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter is the default io.Writer used by cock for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
//		import "github.com/mattn/go-colorable"
//		cock.DefaultWriter = colorable.NewColorableStdout().
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter is the default io.Writer used by cock to debug errors.
var DefaultErrorWriter io.Writer = os.Stderr

var cockMode = debugCode
var modename = DebugMode

// SetMode sets cock mode according to input string.
func SetMode(value string) {
	switch value {
	case DebugMode, "":
		cockMode = debugCode
	case ReleaseMode:
		cockMode = releaseCode
	case TestMode:
		cockMode = testCode
	default:
		panic("cock mode unknown: " + value)
	}

	if value == "" {
		value = DebugMode
	}
	modename = value
}

// Mode returns currently cock mode.
func Mode() string {
	return modename
}

func init() {
	mode := os.Getenv(EnvCockMode)
	SetMode(mode)
}
