package scenes

import "fmt"

type IScene interface {
	GetScene() []string
}

type Scene struct {
	scene []string
}

func (c *Scene) GetScene() []string {
	return c.scene
}

func RenderScene(scene IScene) {
	for i, line := range scene.GetScene() {
		// TODO: Stop hard-coding scene position
		fmt.Printf("\033[%d;30H%s\n", i+3, line)
	}
}
