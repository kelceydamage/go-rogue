package utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type EventTextLoader struct {
	Texts map[string]map[string]map[string][]string `json:"texts"`
}

func NewEventTextLoader() *EventTextLoader {
	return &EventTextLoader{
		Texts: make(map[string]map[string]map[string][]string),
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

func (loader *EventTextLoader) GetText(theme string, nodeType string, subcategory string) string {
	themeTexts, themeExists := loader.Texts[theme]
	if !themeExists {
		return "No text available for this theme."
	}

	nodeTexts, nodeExists := themeTexts[nodeType]
	if !nodeExists {
		return "No text available for this node type."
	}

	// Handle subcategories
	if subcategory != "" {
		subcategoryTexts, subcategoryExists := nodeTexts[subcategory]
		if subcategoryExists && len(subcategoryTexts) > 0 {
			rand.Seed(time.Now().UnixNano())
			return subcategoryTexts[rand.Intn(len(subcategoryTexts))]
		}
		return "No text available for this subcategory."
	}

	// Fallback for non-subcategorized text
	return "No text available."
}

type EdgeTypeScenarios struct {
	Preview    map[string]string `json:"Preview"`
	Text       map[string]string `json:"Text"`
	Transition map[string]string `json:"Transition"`
}

type TraversalTextLoader struct {
	Texts map[string]map[string]*EdgeTypeScenarios `json:"texts"`
}

func NewTraversalTextLoader() *TraversalTextLoader {
	return &TraversalTextLoader{
		Texts: make(map[string]map[string]*EdgeTypeScenarios),
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

func (loader *TraversalTextLoader) GetTraversalTextScenarios(theme string, edgeType string) *EdgeTypeScenarios {
	themeTexts := loader.Texts[theme]
	edgeTexts := themeTexts[edgeType]
	return edgeTexts
}

func (loader *TraversalTextLoader) GetTransitionText(theme string, edgeType string, id string) string {
	themeTexts, themeExists := loader.Texts[theme]
	if !themeExists {
		return "No transition text for this theme."
	}
	edgeTexts, edgeExists := themeTexts[edgeType]
	if !edgeExists || edgeTexts.Transition == nil {
		return "No transition text for this edge type."
	}
	return edgeTexts.Transition[id]
}

func (loader *TraversalTextLoader) GetPreviewText(theme string, edgeType string, id string) string {
	themeTexts, themeExists := loader.Texts[theme]
	if !themeExists {
		return "You see something unknown ahead. 1"
	}
	edgeTexts, edgeExists := themeTexts[edgeType]
	if !edgeExists || len(edgeTexts.Preview) == 0 {
		return "You see something unknown ahead. 2"
	}
	return edgeTexts.Preview[id]
}

func (loader *TraversalTextLoader) GetText(theme string, edgeType string, id string) string {
	themeTexts, themeExists := loader.Texts[theme]
	if !themeExists {
		return "You see something unknown ahead. 1"
	}
	edgeTexts, edgeExists := themeTexts[edgeType]
	if !edgeExists || len(edgeTexts.Text) == 0 {
		return "You see something unknown ahead. 2"
	}
	return edgeTexts.Text[id]
}

type ActionOutcome map[string]string

// ActionData contains success and failure outcomes for an action.
type ActionData struct {
	Success ActionOutcome `json:"Success"`
	Failure ActionOutcome `json:"Failure"`
}

// ActionsLoader holds all actions loaded from the JSON file.
type ActionsLoader struct {
	Actions map[string]ActionData `json:"Actions"`
}

func NewActionsLoader() *ActionsLoader {
	return &ActionsLoader{
		Actions: make(map[string]ActionData),
	}
}

func (loader *ActionsLoader) LoadFromFile(filename string) error {
	fmt.Println("Loading actions file:", filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(data, &loader.Actions)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

// GetRandomSuccess returns a random success text for the given action key.
func (loader *ActionsLoader) GetRandomSuccess(actionKey string) string {
	action, exists := loader.Actions[actionKey]
	if !exists || len(action.Success) == 0 {
		return "No success text available."
	}
	keys := make([]string, 0, len(action.Success))
	for k := range action.Success {
		keys = append(keys, k)
	}
	rand.Seed(time.Now().UnixNano())
	return action.Success[keys[rand.Intn(len(keys))]]
}

// GetRandomFailure returns a random failure text for the given action key.
func (loader *ActionsLoader) GetRandomFailure(actionKey string) string {
	action, exists := loader.Actions[actionKey]
	if !exists || len(action.Failure) == 0 {
		return "No failure text available."
	}
	keys := make([]string, 0, len(action.Failure))
	for k := range action.Failure {
		keys = append(keys, k)
	}
	rand.Seed(time.Now().UnixNano())
	return action.Failure[keys[rand.Intn(len(keys))]]
}

/*
loader := utilities.NewActionsLoader()
err := loader.LoadFromFile("actions.json")
if err != nil {
    log.Fatal(err)
}
fmt.Println(loader.GetRandomSuccess("Swim"))
fmt.Println(loader.GetRandomFailure("Bash"))
*/
