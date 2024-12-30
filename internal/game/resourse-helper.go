package game

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	resource "github.com/quasilyte/ebitengine-resource"
)

func (ctx *Context) NewRect(width, height float64) *graphics.Rect {
	return graphics.NewRect(ctx.graphicsCache, width, height)
}

func (ctx *Context) NewLabel(id resource.FontID) *graphics.Label {
	fnt := ctx.Loader.LoadFont(id)
	return graphics.NewLabel(ctx.graphicsCache, fnt.Face)
}

func (ctx *Context) NewSprite(id resource.ImageID) *graphics.Sprite {
	s := graphics.NewSprite(ctx.graphicsCache)
	if id == 0 {
		return s
	}
	img := ctx.Loader.LoadImage(id)
	s.SetImage(img.Data)
	if img.DefaultFrameWidth != 0 || img.DefaultFrameHeight != 0 {
		w, h := s.GetFrameSize()
		if img.DefaultFrameWidth != 0 {
			w = int(img.DefaultFrameWidth)
		}
		if img.DefaultFrameHeight != 0 {
			h = int(img.DefaultFrameHeight)
		}
		s.SetFrameSize(w, h)
	}
	return s
}
