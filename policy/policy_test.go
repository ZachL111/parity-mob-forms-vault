package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 83, Capacity: 101, Latency: 12, Risk: 9, Weight: 13}
	if got := Score(signal); got != 260 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 85, Capacity: 79, Latency: 24, Risk: 23, Weight: 7}
	if got := Score(signal); got != 120 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 83, Capacity: 92, Latency: 25, Risk: 10, Weight: 12}
	if got := Score(signal); got != 203 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
