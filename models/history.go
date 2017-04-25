package models

// History declares how the trader should be implemented.
type History interface {
	SaveEvent(e Event) (History, error)
	TakeAction() (*Action, error)
}
