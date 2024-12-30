package main

import (
	"TestGame/internal/assets"
	"TestGame/internal/controls"
	"TestGame/internal/game"
	"TestGame/internal/scenes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
)

func main() {
	ctx := game.NewContext()
	ctx.Loader = createLoader()
	ctx.WindowWidth = 480
	ctx.WindowHeight = 480
	ctx.Rand.SetSeed(time.Now().Unix())
	g := &game.Game{
		Ctx: ctx,
	}
	g.InputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	ctx.Input = g.InputSystem.NewHandler(0, controls.DefaultKeymap)

	ebiten.SetWindowSize(g.Ctx.WindowWidth, g.Ctx.WindowHeight)
	ebiten.SetWindowTitle("Ebitengine Quest")

	assets.RegisterResources(ctx.Loader)

	game.ChangeScene(ctx, scenes.NewSplashController(ctx))

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}
