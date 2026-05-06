package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 83, Capacity: 101, Latency: 12, Risk: 9, Weight: 13}, wantScore: 260, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 85, Capacity: 79, Latency: 24, Risk: 23, Weight: 7}, wantScore: 120, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 83, Capacity: 92, Latency: 25, Risk: 10, Weight: 12}, wantScore: 203, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
