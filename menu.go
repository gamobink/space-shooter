package main

import (
	"fmt"
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/otraore/space-shooter/gui"
)

var playBtn *gui.Button

type MenuScene struct{}

func (MenuScene) Preload() {
	err := engo.Files.Load("images/button_silver.png", "images/button_gold.png", "fonts/kenvector_future.ttf")
	if err != nil {
		panic(err)
	}
}

func (MenuScene) Setup(w *ecs.World) {

	common.SetBackground(color.White)

	w.AddSystem(&common.RenderSystem{})

	fnt := &common.Font{
		URL:  "fonts/kenvector_future.ttf",
		FG:   color.White,
		Size: 64,
	}

	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	//Retrieve a texture
	texture, err := common.LoadedSprite("images/button_silver.png")
	if err != nil {
		log.Println(err)
	}

	textureClicked, err := common.LoadedSprite("images/button_gold.png")
	if err != nil {
		log.Println(err)
	}

	x := (engo.GameWidth() / 2) - texture.Width()/2
	y := (engo.GameHeight() / 2) - (texture.Height() / 2) - texture.Height()/2

	fmt.Println(texture.Width())

	playBtn = &gui.Button{
		Text:         "Play",
		World:        w,
		Image:        texture,
		ImageClicked: textureClicked,
		Font:         fnt,
		Position:     engo.Point{x, y},
	}

	playBtn.Init()

	playBtn.OnClick(func() {
		engo.SetScene(GameScene{}, true)
	})

	y += texture.Height() + 30

	exitBtn := &gui.Button{
		Text:         "Exit",
		World:        w,
		Image:        texture,
		ImageClicked: textureClicked,
		Font:         fnt,
		Position:     engo.Point{x, y},
	}
	exitBtn.Init()

	exitBtn.OnClick(func() {
		engo.Exit()
	})

}

func (MenuScene) Type() string { return "MenuScene" }
