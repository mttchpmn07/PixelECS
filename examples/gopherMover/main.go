package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"golang.org/x/image/colornames"
)

const (
	WIDTH  = 800
	HEIGHT = 600
)

func createGophers(gopherAssets []string) []*ecs.Entity {
	gophers := []*ecs.Entity{}

	for i, asset := range gopherAssets {
		gopher, err := ecs.NewEntity()
		if err != nil {
			panic(err)
		}

		loc := NewLocation(500, 500)
		err = gopher.Add(loc)
		if err != nil {
			panic(err)
		}

		kin := NewKenetics(300, 3)
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
		sr, err := NewSprite(asset, active)
		if err != nil {
			fmt.Println(asset)
			panic(err)
		}
		err = gopher.Add(sr)
		if err != nil {
			panic(err)
		}

		sp := NewSpriteProperties(0, 1, sr)
		err = gopher.Add(sp)
		if err != nil {
			panic(err)
		}
		sprop, _ := GetSpriteProperties(gopher)
		sprop.Scale = 150 / sprop.Frame.W()

		gophers = append(gophers, gopher)
	}
	return gophers
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sprite Render Test",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
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

	last := time.Now()
	frames := 0
	second := time.Tick(time.Second)

	controlSystem, err := NewKeybaordControlSystem(gophers...)
	if err != nil {
		panic(err)
	}

	renderSystem, err := NewRendererSystem(gophers...)
	if err != nil {
		panic(err)
	}

	rand.Seed(last.UnixNano())
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		err = controlSystem.Update(win, dt)
		if err != nil {
			panic(err)
		}

		win.Clear(colornames.Skyblue)
		//DrawSprites(win)
		renderSystem.Render(win, dt)
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
