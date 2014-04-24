package unit

import (
	"time"
)

/*Run contains data/metadata about a single run*/
type Run struct {
	distance Distance
	time     time.Duration
}

/*
PacePerMile returns a Duration that represents the average minutes/mile of the Run
*/
func (r Run) PacePerMile() time.Duration {
	return time.Duration(time.Duration(r.time.Seconds()/r.distance.Miles()) * time.Second)
}
