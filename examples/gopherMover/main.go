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
		if i == 0 {
			sr, err := NewSpriteRender(asset, true, loc)
			if err != nil {
				fmt.Println(asset)
				panic(err)
			}
			err = gopher.Add(sr)
			if err != nil {
				panic(err)
			}
		} else {
			sr, err := NewSpriteRender(asset, false, loc)
			if err != nil {
				fmt.Println(asset)
				panic(err)
			}
			err = gopher.Add(sr)
			if err != nil {
				panic(err)
			}
		}
		gophers = append(gophers, gopher)
	}

	last := time.Now()
	active := 0
	elapsed := 3.1
	angle := 0.0
	angleVel := 3.0
	frames := 0
	second := time.Tick(time.Second)

	controlSystem, err := NewControlSystem(300.0, gophers...)
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

		angle += angleVel * dt
		elapsed += dt

		if elapsed > 1.5 && false {
			elapsed = 0
			angleVel *= -1.0
			for i, gopher := range gophers {
				location, err := GetLocation(gopher)
				if err != nil {
					panic(err)
				}
				location.Loc = pixel.V(rand.Float64()*WIDTH, rand.Float64()*HEIGHT)

				render, err := GetSpriteRender(gopher)
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
			render, err := GetSpriteRender(gopher)
			if err != nil {
				panic(err)
			}
			render.Transformation = pixel.IM.Scaled(pixel.ZV, 150/render.Bounds.W()).Rotated(pixel.ZV, angle)
		}

		win.Clear(colornames.Whitesmoke)
		DrawSprites(win)
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
