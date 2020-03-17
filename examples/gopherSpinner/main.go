package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mttchpmn07/CustomECS/ecs"
	"github.com/mttchpmn07/CustomECS/systems"
	"golang.org/x/image/colornames"
)

const (
	WIDTH = 800
	HEIGHT = 600
)

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

	gophers := []*ecs.Entity{}
	gopherAssets := []string{
		"assets/hiking.png",
		"assets/party.png",
		"assets/theif.png",
		"assets/slacker.png",
		"assets/nerdy.png",
		"assets/dragon.png",
	}

	for _, asset := range gopherAssets {
		gopher, err := ecs.NewEntity()
		if err != nil {
			panic(err)
		}
		loc := systems.NewLocation(500, 500)
		gopher.Add(loc)
		sr, err := systems.NewSpriteRender(asset, false, loc)
		if err != nil {
			fmt.Println(asset)
			panic(err)
		}
		gopher.Add(sr)
		gophers = append(gophers, gopher)
	}

	last := time.Now()
	active := 0
	elapsed := 3.1
	angle := 0.0
	angleVel := 3.0
	frames := 0
	second := time.Tick(time.Second)

	rand.Seed(last.UnixNano())
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		angle += angleVel * dt 
		elapsed += dt

		if elapsed > 1.5 {
			elapsed = 0
			angleVel *= -1.0
			for i, gopher := range gophers {
				location, err := systems.GetLocation(gopher)
				if err != nil {
					panic(err)
				}
				location.Loc = pixel.V(rand.Float64()*WIDTH, rand.Float64()*HEIGHT)
	
				render, err := systems.GetSpriteRender(gopher)
				if err != nil {
					panic(err)
				}
				if i == active {
					render.Active = true
				} else {
					render.Active = false
				}
			}
			active++
			if active >= len(gophers) {
				active = 0
			}
		}

		for _, gopher := range gophers {
			render, err := systems.GetSpriteRender(gopher)
			if err != nil {
				panic(err)
			}
			render.Transformation = pixel.IM.Scaled(pixel.ZV, 150/render.Bounds.W()).Rotated(pixel.ZV, angle)
		}

		win.Clear(colornames.Whitesmoke)
		systems.DrawSprites(win)
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
