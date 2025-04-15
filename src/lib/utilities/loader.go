package utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type EventTextLoader struct {
	Texts map[string]map[string][]string `json:"texts"` // Nested map for themes and node types
}

func NewEventTextLoader() *EventTextLoader {
	return &EventTextLoader{
		Texts: make(map[string]map[string][]string),
	}
}

func (loader *EventTextLoader) LoadFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(data, &loader.Texts)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

func (loader *EventTextLoader) GetText(theme string, eventType string) string {
	themeTexts, themeExists := loader.Texts[theme]
	if !themeExists {
		return "No text available for this theme."
	}

	textOptions, eventExists := themeTexts[eventType]
	if !eventExists || len(textOptions) == 0 {
		return "No text available for this event type."
	}

	// Randomize selection
	rand.Seed(time.Now().UnixNano())
	return textOptions[rand.Intn(len(textOptions))]
}

type TraversalTextLoader struct {
	Texts map[string]map[string]struct {
		Preview  []string `json:"Preview"`
		Text     []string `json:"Text"`
		Decision []string `json:"Decision"`
	} `json:"texts"`
}

func NewTraversalTextLoader() *TraversalTextLoader {
	return &TraversalTextLoader{
		Texts: make(map[string]map[string]struct {
			Preview  []string `json:"Preview"`
			Text     []string `json:"Text"`
			Decision []string `json:"Decision"`
		}),
	}
}

func (loader *TraversalTextLoader) LoadFromFile(filename string) error {
	fmt.Println("Loading file:", filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(data, &loader.Texts)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

func (loader *TraversalTextLoader) GetPreview(theme string, edgeType string) string {
	themeTexts, themeExists := loader.Texts[theme]
	if !themeExists {
		return "You see something unknown ahead. 1"
	}

	edgeTexts, edgeExists := themeTexts[edgeType]
	if !edgeExists || len(edgeTexts.Preview) == 0 {
		return "You see something unknown ahead. 2"
	}

	// Randomize selection
	rand.Seed(time.Now().UnixNano())
	return edgeTexts.Preview[rand.Intn(len(edgeTexts.Preview))]
}

func (loader *TraversalTextLoader) GetDecisionText(theme string, edgeType string) string {
	themeTexts, themeExists := loader.Texts[theme]
	if !themeExists {
		return "You traverse the unknown."
	}

	edgeTexts, edgeExists := themeTexts[edgeType]
	if !edgeExists || len(edgeTexts.Text) == 0 {
		return "You traverse the unknown."
	}

	// Randomize selection
	rand.Seed(time.Now().UnixNano())
	return edgeTexts.Decision[rand.Intn(len(edgeTexts.Text))]
}

func (loader *TraversalTextLoader) GetText(theme string, edgeType string) string {
	themeTexts, themeExists := loader.Texts[theme]
	if !themeExists {
		return "You traverse the unknown."
	}

	edgeTexts, edgeExists := themeTexts[edgeType]
	if !edgeExists || len(edgeTexts.Text) == 0 {
		return "You traverse the unknown."
	}

	// Randomize selection
	rand.Seed(time.Now().UnixNano())
	return edgeTexts.Text[rand.Intn(len(edgeTexts.Text))]
}
