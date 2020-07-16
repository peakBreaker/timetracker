package main

import (
	log "github.com/sirupsen/logrus"
)

type Tracker struct{}

func (t Tracker) Track() {
	log.Info("Starting timetracking")
}

func (t Tracker) End() {
	log.Info("Ending timetracking")
}

func List() []string {
	log.Info("Listing entries")

	codes := []string{"code1", "code2"}
	return codes
}

func Edit(hashCode string) {
	log.Info("Editing " + hashCode)
}

func EditLast() {
	log.Info("Starting timetracking")
}

func main() {
	t := Tracker{}
	t.Track()
}
