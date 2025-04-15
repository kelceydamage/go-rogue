package actions

import "go-rogue/src/lib/maps"

type CrawlAction struct {
	name       string
	difficulty int
}

func NewCrawlAction(edge *maps.Edge) *CrawlAction {
	return &CrawlAction{
		name:       "Crawl",
		difficulty: edge.GetDifficulty(),
	}
}

func (a *CrawlAction) GetName() string {
	return a.name
}

func (a *CrawlAction) Execute() {
	// Implement the logic to perform a crawl action
}

func (a *CrawlAction) GetText() string {
	return "Crawl action executed"
}
