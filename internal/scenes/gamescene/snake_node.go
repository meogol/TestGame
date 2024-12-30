package gamescene

import (
	"TestGame/internal/assets"
	"TestGame/internal/controls"
	graphics "github.com/quasilyte/ebitengine-graphics"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
)

type SnakeNode struct {
	input           *input.Handler
	sprite          *graphics.Sprite
	pos             gmath.Vec
	orientation     int
	windowXMax      int
	windowYMax      int
	nextNode        *SnakeNode
	isHead          bool
	state           *sceneState // only head item have it
	count           int
	nextOrientation int
}

const (
	down = iota + 1
	up
	left
	right
)

const subImageSize = 48

func newSnakeNode(pos gmath.Vec, state *sceneState) *SnakeNode {
	return &SnakeNode{pos: pos, orientation: up, isHead: true, state: state, nextNode: nil, count: 0}
}

func newTailSnakeNode(pos gmath.Vec, orientation int, c *Controller) *SnakeNode {
	var newPos gmath.Vec
	if orientation == up {
		newPos = gmath.Vec{X: pos.X, Y: pos.Y + subImageSize}
	} else if orientation == down {
		newPos = gmath.Vec{X: pos.X, Y: pos.Y - subImageSize}
	} else if orientation == left {
		newPos = gmath.Vec{X: pos.X + subImageSize, Y: pos.Y}
	} else if orientation == right {
		newPos = gmath.Vec{X: pos.X - subImageSize, Y: pos.Y}
	}

	return &SnakeNode{pos: newPos, orientation: orientation, isHead: false, nextNode: c.state.tailItem}
}

func (sn *SnakeNode) Init(s *scene) {
	ctx := s.Controller().ctx

	sn.windowXMax = s.Controller().ctx.WindowWidth
	sn.windowYMax = s.Controller().ctx.WindowHeight
	sn.input = ctx.Input

	sprite := s.Controller().ctx.NewSprite(assets.Character)

	sprite.SetFrameSize(subImageSize, subImageSize)
	sprite.SetFrameOffset(0, sn.orientation*subImageSize)

	sprite.Pos.Base = &sn.pos

	sn.sprite = sprite
	s.AddGraphics(sn.sprite)
}

func (sn *SnakeNode) IsDisposed() bool {
	return false
}

func (sn *SnakeNode) Update(delta float64) {
	if sn.isHead && sn.count == sn.state.frameDelay {
		sn.count = 0
		moveTail(sn.state.tailItem)
		offset := sn.move()
		sn.pos = sn.pos.Add(offset)

		if sn.pos.X <= 0 {
			sn.pos.X = float64(sn.windowXMax)
		} else if sn.pos.X > float64(sn.windowXMax) {
			sn.pos.X = 0
		}

		if sn.pos.Y < 0 {
			sn.pos.Y = float64(sn.windowYMax)
		} else if sn.pos.Y > float64(sn.windowYMax) {
			sn.pos.Y = 0
		}
		sn.nextOrientation = 0
	} else {
		if sn.isHead {
			if sn.input.ActionIsPressed(controls.ActionMoveRight) {
				sn.nextOrientation = right
			} else if sn.input.ActionIsPressed(controls.ActionMoveDown) {
				sn.nextOrientation = down
			} else if sn.input.ActionIsPressed(controls.ActionMoveLeft) {
				sn.nextOrientation = left
			} else if sn.input.ActionIsPressed(controls.ActionMoveUp) {
				sn.nextOrientation = up
			}
		}
		sn.count++
	}
}

func (sn *SnakeNode) move() gmath.Vec {
	speed := float64(subImageSize)
	var v gmath.Vec

	if sn.nextOrientation != 0 {
		sn.orientation = sn.nextOrientation
	}

	if sn.input.ActionIsPressed(controls.ActionMoveRight) {
		sn.orientation = right
		v.X += speed
	} else if sn.input.ActionIsPressed(controls.ActionMoveDown) {
		sn.orientation = down
		v.Y += speed
	} else if sn.input.ActionIsPressed(controls.ActionMoveLeft) {
		sn.orientation = left
		v.X -= speed
	} else if sn.input.ActionIsPressed(controls.ActionMoveUp) {
		sn.orientation = up
		v.Y -= speed
	} else {
		if sn.orientation == right {
			v.X += speed
		} else if sn.orientation == down {
			v.Y += speed
		} else if sn.orientation == left {
			v.X -= speed
		} else if sn.orientation == up {
			v.Y -= speed
		}
	}

	return v
}

func moveTail(tailNode *SnakeNode) {
	if !(tailNode.isHead) && tailNode.nextNode != nil {
		tailNode.pos = tailNode.nextNode.pos
		moveTail(tailNode.nextNode)
	}
}
