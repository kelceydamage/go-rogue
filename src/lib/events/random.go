package events

type NarrativeEvent struct {
	text string
}

func (e *NarrativeEvent) Execute() {}

func (e *NarrativeEvent) GetText() string {
	return e.text
}
