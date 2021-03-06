package anim

import (
	"korok.io/korok/anim/frame"
	"korok.io/korok/anim/tween"
	"korok.io/korok/gfx"
)

type AnimationSystem struct {
	SpriteEngine *frame.Engine
	TweenEngine *tween.Engine

	// tables
	st *gfx.SpriteTable
	xf *gfx.TransformTable
}

func NewAnimationSystem() *AnimationSystem {
	var (
		se = frame.NewEngine()
		te = tween.NewEngine()
	)
	spriteEngine = se
	tweenEngine = te
	as = &AnimationSystem{
		SpriteEngine:se,
		TweenEngine:te,
	}
	return as
}

func (as *AnimationSystem) RequireTable(tables []interface{}) {
	as.SpriteEngine.RequireTable(tables)

	for _, t := range tables {
		switch table := t.(type) {
		case *gfx.SpriteTable:
			as.st = table
		case *gfx.TransformTable:
			as.xf = table
		}
	}
}

func (as *AnimationSystem) Update(dt float32) {
	as.SpriteEngine.Update(dt)
	as.TweenEngine.Update(dt)
}

// shortcut
var spriteEngine *frame.Engine
var tweenEngine *tween.Engine
var as *AnimationSystem