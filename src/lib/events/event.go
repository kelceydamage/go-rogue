package events

type Event interface {
	Execute()
	GetText() string
}
