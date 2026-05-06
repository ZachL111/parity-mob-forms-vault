package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 58, Slack: 28, Drag: 24, Confidence: 67}
	if got := DomainReviewScore(item); got != 139 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "watch" {
		t.Fatalf("lane = %s", got)
	}
}
