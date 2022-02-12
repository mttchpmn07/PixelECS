package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/mttchpmn07/PixelECS/shapePlacer/entities"
	"github.com/mttchpmn07/PixelECS/shapePlacer/systems"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	width  = 800
	height = 600
)

func buildSystems() {
	// Batch renderer system
	renderSystem, err := systems.NewSRender()
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(renderSystem)
	if err != nil {
		panic(err)
	}

	square, err := entities.NewSquare(collision2d.NewVector(width/2, height/2), 0, 100)
	if err != nil {
		panic(err)
	}

	err = renderSystem.AddEntity(square)
	if err != nil {
		panic(err)
	}
	// User Control System
	controlSystem, err := systems.NewSUsuerControl()
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(controlSystem)
	if err != nil {
		panic(err)
	}
}

func updateFPS(win *pixelgl.Window, cfg pixelgl.WindowConfig, frames int, second <-chan time.Time) {
	select {
	case <-second:
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
	default:
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Shape Placer",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	buildSystems()

	frames := 0
	second := time.Tick(time.Second)
	last := time.Now()
	rand.Seed(last.UnixNano())

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Skyblue)

		err = ecs.UpdateSystems(win, &dt)
		if err != nil {
			panic(err)
		}
		win.Update()

		updateFPS(win, cfg, frames, second)
		frames++
	}
}

func main() {
	pixelgl.Run(run)
}
