package events

type EncounterEvent struct {
	text string
}

func (e *EncounterEvent) Execute() {}

func (e *EncounterEvent) GetText() string {
	return e.text
}
