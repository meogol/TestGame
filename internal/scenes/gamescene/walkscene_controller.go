package gamescene

import (
	"TestGame/internal/controls"
	"TestGame/internal/game"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"log"
	"math"
)

type scene = gscene.Scene[*Controller]

type Controller struct {
	ctx *game.Context

	scene *gscene.RootScene[*Controller]
	state *sceneState

	scoreLabel *graphics.Label
	score      int
}

func NewController(ctx *game.Context) *Controller {
	return &Controller{ctx: ctx}
}

func (c *Controller) Init(s *gscene.RootScene[*Controller]) {
	c.scene = s

	c.state = &sceneState{frameDelay: 48}
	sn := newSnakeNode(gmath.Vec{
		X: math.Round(c.ctx.Rand.FloatRange(0, float64(c.ctx.WindowWidth))/48) * 48,
		Y: math.Round(c.ctx.Rand.FloatRange(0, float64(c.ctx.WindowHeight))/48) * 48,
	}, c.state)
	c.state.headItem = sn
	c.state.AddSnakeNode(sn)

	s.AddObject(sn)
	c.createPickup()
}

func (c *Controller) createPickup() {
	p := newPickupNode(gmath.Vec{
		X: math.Round(c.ctx.Rand.FloatRange(0, float64(c.ctx.WindowWidth))/48) * 48,
		Y: math.Round(c.ctx.Rand.FloatRange(0, float64(c.ctx.WindowHeight))/48) * 48,
	})
	p.EventDestroyed.Connect(nil, func(score int) {
		c.createPickup()

		tail := c.state.tailItem
		orientation := tail.orientation
		if tail.nextOrientation != 0 {
			orientation = tail.nextOrientation
		}
		sn := newTailSnakeNode(tail.pos, orientation, c)
		c.state.AddSnakeNode(sn)
		c.scene.AddObject(sn)
		log.Println("add new snake node")
	})

	c.scene.AddObject(p)
}

func (c *Controller) Update(delta float64) {
	if c.ctx.Input.ActionIsJustPressed(controls.ActionRestart) {
		game.ChangeScene(c.ctx, NewController(c.ctx))
	}
}
