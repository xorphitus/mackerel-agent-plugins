package mpunicorn

import "testing"

func TestFetchMetrics(t *testing.T) {
	var unicorn UnicornPlugin

	stat, _ := unicorn.FetchMetrics()
	if len(stat) != 7 {
		t.Errorf("GetStat: %d should be 7", len(stat))
	}
}

func TestGraphDefinition(t *testing.T) {
	var unicorn UnicornPlugin

	graphdef := unicorn.GraphDefinition()
	if len(graphdef) != 3 {
		t.Errorf("GetTempfilename: %d should be 3", len(graphdef))
	}
}
