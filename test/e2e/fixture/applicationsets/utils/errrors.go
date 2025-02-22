package utils

import (
	"os"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

const (
	// ErrorCommandSpecific is reserved for command specific indications
	ErrorCommandSpecific = 1
	// ErrorConnectionFailure is returned on connection failure to API endpoint
	ErrorConnectionFailure = 11
	// ErrorAPIResponse is returned on unexpected API response, i.e. authorization failure
	ErrorAPIResponse = 12
	// ErrorResourceDoesNotExist is returned when the requested resource does not exist
	ErrorResourceDoesNotExist = 13
	// ErrorGeneric is returned for generic error
	ErrorGeneric = 20
)

// CheckError logs a fatal message and exits with ErrorGeneric if err is not nil
func CheckError(err error) {
	if err != nil {
		debug.PrintStack()
		Fatal(ErrorGeneric, err)
	}
}

// CheckErrorWithCode is a convenience function to exit if an error is non-nil and exit if it was
func CheckErrorWithCode(err error, exitcode int) {
	if err != nil {
		Fatal(exitcode, err)
	}
}

// FailOnErr panics if there is an error. It returns the first value so you can use it if you cast it:
// text := FailOrErr(Foo)).(string)
func FailOnErr(v any, err error) any {
	CheckError(err)
	return v
}

// Fatal is a wrapper for logrus.Fatal() to exit with custom code
func Fatal(exitcode int, args ...any) {
	exitfunc := func() {
		os.Exit(exitcode)
	}
	log.RegisterExitHandler(exitfunc)
	log.Fatal(args...)
}

// Fatalf is a wrapper for logrus.Fatalf() to exit with custom code
func Fatalf(exitcode int, format string, args ...any) {
	exitfunc := func() {
		os.Exit(exitcode)
	}
	log.RegisterExitHandler(exitfunc)
	log.Fatalf(format, args...)
}
