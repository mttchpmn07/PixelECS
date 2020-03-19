package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/entities"
	"github.com/mttchpmn07/PixelECS/gopherMover/systems"
	"golang.org/x/image/colornames"
)

const (
	width  = 800
	height = 600
)

func createGophers(gopherAssets []string) []*ecs.Entity {
	gophers := []*ecs.Entity{}

	for _, asset := range gopherAssets {
		gopher, err := entities.NewGopher(asset, width/2, height/2)
		if err != nil {
			panic(err)
		}
		gophers = append(gophers, gopher)
	}
	return gophers
}

func buildSystems(gophers []*ecs.Entity) {
	controlSystem, err := systems.NewSKeyboardController(gophers...)
	if err != nil {
		panic(err)
	}
	ecs.RegisterSystem(controlSystem)

	renderSystem, err := systems.NewSRenderer(gophers...)
	if err != nil {
		panic(err)
	}
	ecs.RegisterSystem(renderSystem)
}

func buildWindow() (pixelgl.WindowConfig, *pixelgl.Window) {
	cfg := pixelgl.WindowConfig{
		Title:  "Sprite Render Test",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)
	return cfg, win
}

func updateFrames(win *pixelgl.Window, cfg pixelgl.WindowConfig, frames int, second <-chan time.Time) {
	select {
	case <-second:
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
	default:
	}
}

func run() {
	cfg, win := buildWindow()
	gopherAssets := []string{
		//"assets/hiking.png",
		//"assets/party.png",
		//"assets/theif.png",
		//"assets/slacker.png",
		//"assets/nerdy.png",
		"assets/dragon.png",
	}
	gophers := createGophers(gopherAssets)
	buildSystems(gophers)

	frames := 0
	second := time.Tick(time.Second)
	last := time.Now()
	rand.Seed(last.UnixNano())
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Skyblue)
		ecs.UpdateSystems(win, &dt)
		win.Update()

		updateFrames(win, cfg, frames, second)
		frames++
	}
}

func main() {
	pixelgl.Run(run)
}
