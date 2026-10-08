package pilot

import "testing"

func TestNormalizeCandidate(t *testing.T) {
	if got := NormalizeCandidate("  Merge-Group-01  "); got != "merge-group-01" {
		t.Fatalf("NormalizeCandidate() = %q", got)
	}
}

func TestNormalizeCandidateEmpty(t *testing.T) {
	if got := NormalizeCandidate("  "); got != "" {
		t.Fatalf("NormalizeCandidate() = %q, want empty", got)
	}
}

func TestGovernanceSample(t *testing.T) {
	if got := GovernanceSample(); got != "positive" {
		t.Fatalf("GovernanceSample() = %q", got)
	}
}
