package main

import (
	"fmt"
	"image"
	"math/rand"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherPlatformer/components"
	"github.com/mttchpmn07/PixelECS/gopherPlatformer/entities"
	"github.com/mttchpmn07/PixelECS/gopherPlatformer/systems"
	"golang.org/x/image/colornames"
)

const (
	width  = 800
	height = 600
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func createWalls() []*ecs.Entity {
	walls := []*ecs.Entity{}

	// right wall
	wall, err := entities.NewWall(width-5, height/2, 10, height)
	if err != nil {
		panic(err)
	}
	walls = append(walls, wall)

	// left wall
	wall, err = entities.NewWall(5, height/2, 10, height)
	if err != nil {
		panic(err)
	}
	walls = append(walls, wall)

	// bottom wall
	wall, err = entities.NewWall(width/2, 5, width, 10)
	if err != nil {
		panic(err)
	}
	walls = append(walls, wall)

	// top wall
	wall, err = entities.NewWall(width/2, height-5, width, 10)
	if err != nil {
		panic(err)
	}
	walls = append(walls, wall)

	return walls
}

func createGophers(gopherAsset string) *ecs.Entity {
	ba, err := components.NewCBatchAsset(gopherAsset)
	if err != nil {
		panic(err)
	}

	gopher, err := entities.NewGopher(width, height, 100, ba.(*components.CBatchAsset))
	if err != nil {
		panic(err)
	}
	return gopher
}

func createFlys(num int, flyAsset string) []*ecs.Entity {
	flys := []*ecs.Entity{}
	ba, err := components.NewCBatchAsset(flyAsset)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= num; i++ {
		fly, err := entities.NewFly(width, height, 25, ba.(*components.CBatchAsset))
		if err != nil {
			panic(err)
		}
		flys = append(flys, fly)
	}
	return flys
}

func buildSystems(gopher *ecs.Entity, flys []*ecs.Entity, walls []*ecs.Entity) {
	// Random Walk System
	randoWalkSystem, err := systems.NewSRandomWalk(flys...)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(randoWalkSystem)
	if err != nil {
		panic(err)
	}

	// Animated Sprite Render System
	animatorSystem, err := systems.NewSAnimator(flys...)
	if err != nil {
		panic(err)
	}
	err = animatorSystem.AddEntity(gopher)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(animatorSystem)
	if err != nil {
		panic(err)
	}

	// Batch renderer system
	batchRendererSystem, err := systems.NewSBatchRenderer(flys...)
	if err != nil {
		panic(err)
	}
	err = batchRendererSystem.AddEntity(gopher)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(batchRendererSystem)
	if err != nil {
		panic(err)
	}

	// Keyboard Control System
	controlSystem, err := systems.NewSKeyboardController(gopher)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(controlSystem)
	if err != nil {
		panic(err)
	}

	collisionSystem, err := systems.NewSCollisionTracker(width, height, flys...)
	if err != nil {
		panic(err)
	}
	err = collisionSystem.AddEntity(walls...)
	if err != nil {
		panic(err)
	}
	err = collisionSystem.AddEntity(gopher)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(collisionSystem)
	if err != nil {
		panic(err)
	}

	collisionRenderSystem, err := systems.NewSCollisionPolyRenderer(flys...)
	if err != nil {
		panic(err)
	}
	err = collisionRenderSystem.AddEntity(gopher)
	if err != nil {
		panic(err)
	}
	err = collisionRenderSystem.AddEntity(walls...)
	if err != nil {
		panic(err)
	}
	err = ecs.RegisterSystem(collisionRenderSystem)
	if err != nil {
		panic(err)
	}

	physicsSystem, err := systems.NewSPhysics(flys...)
	if err != nil {
		panic(err)
	}
	/*
		err = collisionRenderSystem.AddEntity(gopher)
		if err != nil {
			panic(err)
		}
		err = collisionRenderSystem.AddEntity(walls...)
		if err != nil {
			panic(err)
		}
	*/
	err = ecs.RegisterSystem(physicsSystem)
	if err != nil {
		panic(err)
	}
}

func buildWindow() (pixelgl.WindowConfig, *pixelgl.Window) {
	cfg := pixelgl.WindowConfig{
		Title:  "Sprite Render Test",
		Bounds: pixel.R(0, 0, width, height),
		//VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	//win.SetSmooth(true)
	return cfg, win
}

func updateFPS(win *pixelgl.Window, cfg pixelgl.WindowConfig, frames int, second <-chan time.Time) {
	select {
	case <-second:
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
	default:
	}
}

func run() {
	//cfg, win := buildWindow()
	cfg := pixelgl.WindowConfig{
		Title:  "GopherPlatformer",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	walls := createWalls()
	gopher := createGophers("assets/dragon_animated.png")
	flys := createFlys(300, "assets/bug.png")
	buildSystems(gopher, flys, walls)

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
