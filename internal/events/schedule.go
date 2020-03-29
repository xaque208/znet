package events

import (
	"sort"
	"time"

	"github.com/apex/log"
)

// Scheduler holds timeSlice objects and provides an methods to update them..
type Scheduler struct {
	timeSlice *TimeSlice
}

// NewScheduler returns a new Scheduler.
func NewScheduler() *Scheduler {
	return &Scheduler{
		timeSlice: &TimeSlice{},
	}
}

// All returns all current timeSlice objects.
func (s *Scheduler) All() TimeSlice {
	return *s.timeSlice
}

// Next determines the next occurring event in the series.
func (s *Scheduler) Next() time.Time {

	keys := []time.Time{}

	for k := range *s.timeSlice {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	return keys[0]
}

// NamesForTime returns all events names that are scheduled for a given timeSlice.
func (s *Scheduler) NamesForTime(t time.Time) []string {
	return (*s.timeSlice)[t]
}

// WaitForNext is a blocking function that waits for the next available time to
// arrive before returning the names to the caller.
func (s *Scheduler) WaitForNext() []string {
	next := s.Next()

	if time.Now().After(next) {
		log.Infof("sending past event: %s", next)
		return s.NamesForTime(next)
	}

	log.Infof("waiting until: %s", next)
	ti := time.NewTimer(time.Until(next))
	<-ti.C

	return s.NamesForTime(next)
}

// Step deletes the next timeslice.  This is determined to be the timeslice
// that has just run.  The expectataion is that Step() is called once the
// events have completed firing to advance to the next position in time.
func (s *Scheduler) Step() {
	delete(*s.timeSlice, s.Next())
}

// Set appends the name given to the time slot given.
func (s *Scheduler) Set(t time.Time, name string) {

	if _, ok := (*s.timeSlice)[t]; !ok {
		(*s.timeSlice)[t] = make([]string, 1)
	}

	timeHasName := func(names []string) bool {
		for _, n := range names {
			if n == name {
				return true
			}
		}

		return false
	}((*s.timeSlice)[t])

	if !timeHasName {
		(*s.timeSlice)[t] = append((*s.timeSlice)[t], name)
	}

}

// TimeSlice is an association between a specific time, and the names of the events that should fire at that time.
type TimeSlice map[time.Time][]string
