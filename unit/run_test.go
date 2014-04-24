package unit

import (
	"testing"
	"time"
)

func TestPacePerMile(t *testing.T) {
	run := Run{distance: Miles(3), time: time.Duration(21 * time.Minute)}
	pace := run.PacePerMile()

	if pace != time.Duration(7*time.Minute) {
		t.Errorf("Incorrect pace conversion", run.PacePerMile())
	}
}
