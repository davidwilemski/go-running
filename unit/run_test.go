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

func TestPacePerKilometer(t *testing.T) {
	run := Run{distance: Kilometers(5), time: time.Duration(20 * time.Minute)}
	pace := run.PacePerKilometer()

	if pace != time.Duration(4*time.Minute) {
		t.Errorf("Incorrect pace conversion", run.PacePerKilometer())
	}

	run = Run{distance: Kilometers(5), time: time.Duration(25 * time.Minute)}
	pace = run.PacePerKilometer()

	if pace != time.Duration(5*time.Minute) {
		t.Errorf("Incorrect pace conversion", run.PacePerKilometer())
	}
}
