package pilot

import "strings"

// NormalizeCandidate returns the canonical identifier used by the pilot fixture.
func NormalizeCandidate(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

// GovernanceSample identifies the live positive governance sample.
func GovernanceSample() string { return "positive" }
