package components

import (
	"fmt"
	"go-rogue/src/lib/actions"
	"go-rogue/src/lib/maps"
)

func PresentEdgeActions(edge *maps.Edge) []actions.Action {
	var availableActions []actions.Action

	// Map edge types to relevant actions
	switch edge.GetMetaData().Name {
	case maps.Path:
		availableActions = []actions.Action{
			actions.NewProceedAction(edge),
		}
	case maps.Crossing:
		availableActions = []actions.Action{
			actions.NewSwimAction(edge),
			actions.NewJumpAction(edge),
		}
	case maps.Tunnel:
		availableActions = []actions.Action{
			actions.NewCrawlAction(edge),
			actions.NewLightTorchAction(edge),
			actions.NewProceedAction(edge),
		}
	case maps.UnlockedDoor:
		availableActions = []actions.Action{
			actions.NewOpenAction(edge),
		}
	case maps.LockedDoor:
		availableActions = []actions.Action{
			actions.NewBashAction(edge),
			actions.NewLockPickAction(edge),
			actions.NewUnlockAction(edge),
		}
	case maps.HiddenDoor:
		availableActions = []actions.Action{
			actions.NewInvestigateAction(edge),
		}
	default:
		fmt.Println("You encounter an unknown obstacle.")
	}
	return availableActions
}
