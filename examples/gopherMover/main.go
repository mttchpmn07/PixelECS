package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
	"github.com/mttchpmn07/PixelECS/gopherMover/systems"
	"golang.org/x/image/colornames"
)

const (
	width  = 800
	height = 600
)

func createGophers(gopherAssets []string) []*ecs.Entity {
	gophers := []*ecs.Entity{}

	for i, asset := range gopherAssets {
		gopher, err := ecs.NewEntity()
		if err != nil {
			panic(err)
		}

		loc := components.NewCLocation(500, 500)
		err = gopher.Add(loc)
		if err != nil {
			panic(err)
		}

		kin := components.NewCKenetics(300, 3)
		err = gopher.Add(kin)
		if err != nil {
			panic(err)
		}

		var active bool
		if i == 0 {
			active = true
		} else {
			active = false
		}
		sr, err := components.NewCSprite(asset, active)
		if err != nil {
			fmt.Println(asset)
			panic(err)
		}
		err = gopher.Add(sr)
		if err != nil {
			panic(err)
		}

		sp := components.NewCSpriteProperties(0, 1, sr)
		err = gopher.Add(sp)
		if err != nil {
			panic(err)
		}
		sprop, _ := components.GetCSpriteProperties(gopher)
		sprop.Scale = 150 / sprop.Frame.W()

		gophers = append(gophers, gopher)
	}
	return gophers
}

func run() {
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

	gopherAssets := []string{
		"assets/hiking.png",
		"assets/party.png",
		"assets/theif.png",
		"assets/slacker.png",
		"assets/nerdy.png",
		"assets/dragon.png",
	}
	gophers := createGophers(gopherAssets)

	controlSystem, err := systems.NewSKeyboardController(gophers...)
	if err != nil {
		panic(err)
	}

	renderSystem, err := systems.NewSRenderer(gophers...)
	if err != nil {
		panic(err)
	}

	systems := []ecs.System{
		controlSystem,
		renderSystem,
	}

	frames := 0
	second := time.Tick(time.Second)
	last := time.Now()
	rand.Seed(last.UnixNano())
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Skyblue)
		for _, sys := range systems {
			err := sys.Update(win, dt)
			if err != nil {
				panic(err)
			}
		}
		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
