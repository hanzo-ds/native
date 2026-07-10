package tracing

import "github.com/hanzo-ds/native/internal/version"

// Version is the current release version of the ch instrumentation.
func Version() string {
	return version.Get().Raw
}

// SemVersion is the semantic version to be supplied to tracer/meter creation.
func SemVersion() string {
	return "semver:" + Version()
}
