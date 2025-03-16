package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
	"todo/frontend"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS
var store string
var Todolist []string
func GetFiber() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE",
	}))


	app.Post("/add",func(c*fiber.Ctx)error{
		task:=c.FormValue("task")
		store +=" "+task
		Todolist = strings.Split(store," ")
		fmt.Println(store,Todolist)

		return c.SendString("Added")
	})
   
	app.Get("/render", func(c *fiber.Ctx) error {

		component := frontend.Todo(Todolist)
		return component.Render(c.Context(), c)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		indexStr := c.Query("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return err
		}
		if index >= 0 && index < len(Todolist) {
			Todolist = append(Todolist[:index], Todolist[index+1:]...)
			store = strings.Join(Todolist, " ")
		}
		fmt.Println(store, Todolist)
		component := frontend.Todo(Todolist)
		return component.Render(c.Context(), c)
	})

	fmt.Println(store,Todolist)
	app.Listen(":3000")
}


func main() {
	go GetFiber()
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "TODO",
		Width:  800,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent: true,
			
			BackdropType:windows.None,
			DisableFramelessWindowDecorations:false,
		},
	
	OnStartup:        app.startup,
		
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
