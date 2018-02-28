package main

import (
        "log"
	"engo.io/engo"
	"image/color"
	"engo.io/engo/common"
	"engo.io/ecs"
)

type myScene struct {}

// Type uniquely defines your game type
func (*myScene) Type() string { return "goGoCloud" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
  engo.Files.Load("textures/cumulus.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(world *ecs.World) {
  world.AddSystem(&common.RenderSystem{})
  cloud := Cloud{BasicEntity: ecs.NewBasic()}

  //define the position by providing data to the spacecomponent
  cloud.SpaceComponent = common.SpaceComponent{
    Position: engo.Point{10, 10},
    Width:    303,
    Height:   641,
  }

  //load the texture
  texture, err := common.LoadedSprite("textures/cumulus.png")
  if err != nil {
      log.Println("Unable to load texture: " + err.Error())
  }

  //attempt to render the cloud component
  cloud.RenderComponent = common.RenderComponent{
      Drawable: texture,
      Scale:    engo.Point{1, 1},
  }

  //Add the component to the system
  // we loop over components and insert our system
  for _, system := range world.Systems() {
    switch sys := system.(type) {
    case *common.RenderSystem:
        sys.Add(&cloud.BasicEntity, &cloud.RenderComponent, &cloud.SpaceComponent)
    }
  }
  common.SetBackground(color.White)
}

type Cloud struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func main() {
	opts := engo.RunOptions{
		Title: "Hello World",
		Width:  400,
		Height: 400,
	}
	engo.Run(opts, &myScene{})
}
