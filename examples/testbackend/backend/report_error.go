package backend

import (
	// "fmt"
	// "github.com/getsentry/sentry-go"
	"log"
)

func reportError(message string, err error) {
	// sentry.CaptureException(fmt.Errorf("%s: %v", message, err))
	log.Printf("🚨 %s – %v", message, err)
}
