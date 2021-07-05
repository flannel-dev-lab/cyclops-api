package cycapi

import (
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)

var SentryEnabled = false

func Init(dsn, environment string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         dsn,
		Environment: environment,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	SentryEnabled = true
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
}

func Error(err error) {
	if SentryEnabled {
		sentry.CaptureException(err)
	}
	println("Error: " + err.Error())
}
