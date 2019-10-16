package main

import (
	"fmt"

	"github.com/leaanthony/mewn"
	tuxedokeyboard "github.com/leaanthony/tuxedo-keyboard"
	"github.com/wailsapp/wails"
)

func updateKeyboardLights(r, g, b uint8, a float32) {
	// We need to convert the alpha value (0-1)
	// to brightness (0-255)
	brightness := uint8(255 * a)

	// Set the brightness
	err := tuxedokeyboard.SetBrightness(brightness)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Set the Keyboard colour left
	err = tuxedokeyboard.SetKeyboardColourLeft(r, g, b)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Set the Keyboard colour center
	err = tuxedokeyboard.SetKeyboardColourCenter(r, g, b)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Set the Keyboard colour right
	err = tuxedokeyboard.SetKeyboardColourRight(r, g, b)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  220,
		Height: 310,
		Title:  "Tux Key UI",
		JS:     js,
		CSS:    css,
		Colour: "#000000",
	})
	app.Bind(updateKeyboardLights)
	app.Run()
}
