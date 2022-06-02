package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// logNetworkRequestStdOut adds a logger wrapper to add extra network client information to the log
//func logNeworkRequestStdOut(s string, r *http.Request) {
//	logStdOut("IP[" + ReadUserIP(r) + "] UA[" + r.UserAgent() + "] " + string(s))
//}

// logStdOut just logs something to stdout
func logStdOut(s string) {
	log.Printf("%s\n", string(s))
}

// logStdErr just logs to stderr
func logStdErr(s string) {
	log.Fatalf("%s\n", string(s))
}

// Stoerr wraps a string in an error object
func Stoerr(s string) error {
	return &errorString{s}
}

// logStringAsErrorToStdout wraps a string in an error object and pushes to stdout
//func logStringAsErrorToStdout(s string) {
//	check(&errorString{s})
//}

// logStringAsErrorToStderr wraps a string in an error object and pushes to stdout
//func logStringAsErrorToStderr(s string) {
//	checkAndFail(&errorString{s})
//}

// check does error checking
func check(e error) {
	if e != nil {
		log.Printf("error: %v", e)
	}
}

// checkAndFail checks for an error type and fails
func checkAndFail(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	// We wrap our anonymous function, and cast it to a http.HandlerFunc
	// Because our function signature matches ServeHTTP(w, r), this allows
	// our function (type) to implicitly satisify the http.Handler interface.
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Logic before - reading request values, putting things into the
			// request context, performing authentication
			start := time.Now()

			// Important that we call the 'next' handler in the chain. If we don't,
			// then request handling will stop here.
			next.ServeHTTP(w, r)

			// Logic after - useful for logging, metrics, etc.

			ri := &HTTPReqInfo{
				method:    r.Method,
				uri:       r.URL.String(),
				referer:   r.Header.Get("Referer"),
				userAgent: r.Header.Get("User-Agent"),
				ipaddr:    ReadUserIP(r),
				code:      http.StatusOK,
				duration:  time.Since(start),
			}

			logStdOut("IP[" + ri.ipaddr + "] UA[" + ri.userAgent + "] " + ri.method + ":" + fmt.Sprint(ri.code) + " " + ri.uri + " - took " + fmt.Sprint(ri.duration.Seconds()) + "s")
			// It's important that we don't use the ResponseWriter after we've called the
			// next handler: we may cause conflicts when trying to write the response
		},
	)
}
