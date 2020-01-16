package game

import (
	"backend/objects"
	"time"
)

type Game struct {
	Width  int
	Height int

	IsStarted bool
	StartTime time.Time

	ship *objects.Ship
}
