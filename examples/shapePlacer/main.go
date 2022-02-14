package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/mttchpmn07/PixelECS/shapePlacer/entities"
	"github.com/mttchpmn07/PixelECS/shapePlacer/systems"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
)

const (
	width  = 800
	height = 600
)

func buildWindow(title string) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)
	return win
}

func buildSystems() {
	// Create messenger
	m := messenger.NewMessenger()

	// Camera entity
	camera, err := entities.NewCamera(pixel.ZV, 500.0, 1.0, 1.2)
	if err != nil {
		panic(err)
	}

	// Camera system
	cameraSystem, err := systems.NewSCamera(m, camera)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(cameraSystem)
	if err != nil {
		panic(err)
	}

	// User Input System
	userInputSystem, err := systems.NewSUserInput(m, camera)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(userInputSystem)
	if err != nil {
		panic(err)
	}

	// Shape Placer System
	shapePlacerSystem, err := systems.NewSShapePlacer(m)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(shapePlacerSystem)
	if err != nil {
		panic(err)
	}

	// Batch renderer system
	renderSystem, err := systems.NewSRender(m)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(renderSystem)
	if err != nil {
		panic(err)
	}
}

func updateFPS(win *pixelgl.Window, title string, frames int, second <-chan time.Time) {
	select {
	case <-second:
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", title, frames))
	default:
	}
}

func run() {
	title := "Shape Placer"

	win := buildWindow(title)
	buildSystems()

	frames := 0
	second := time.Tick(time.Second)
	last := time.Now()
	rand.Seed(last.UnixNano())

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Skyblue)

		err := ecs.UpdateSystems(win, &dt)
		if err != nil {
			panic(err)
		}
		win.Update()

		updateFPS(win, title, frames, second)
		frames++
	}
}

func main() {
	pixelgl.Run(run)
}
