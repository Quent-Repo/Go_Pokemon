package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	image_url := "PokemonData/Abra/6119be89c1d24539a0310f40b947afc1.jpg"
	q := []string{}
	v := []string{}

	a := app.New()
	w := a.NewWindow("This is a window")
	files, err := ioutil.ReadDir("PokemonData/Arbok")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		v = append(v, file.Name())
	}

	//image from online
	sun, _ := fyne.LoadResourceFromURLString("https://assets.pokemon.com/assets/cms2/img/pokedex/full/181.png")
	sun2 := canvas.NewImageFromResource(sun)

	image := canvas.NewImageFromFile(image_url)
	//image_url2 := "PokemonData/Abra/b0b6de31451f4e7aa3411fe0963a7f4f.jpg"
	//image2 := canvas.NewImageFromFile(image_url2)
	count := 1
	text1 := canvas.NewText(strconv.Itoa(count)+"/150", color.Black)
	text2 := canvas.NewText("one", color.White)
	text3 := canvas.NewText("one", color.White)
	Test2 := widget.NewButton("Keep", func() {
		image.File = image_url
		image.Refresh()
		image.Resize(fyne.NewSize(300, 300))
		fmt.Println(image.Size(), image.File)
		q = append(q, "Keep")
		count++
		fmt.Println(count)
		text1.Text = strconv.Itoa(count) + "/150"
		text1.Refresh()
	})
	Test3 := widget.NewButton("Throw", func() {
		image.File = "PokemonData/Arbok/" + v[count]
		image.Refresh()
		image.Resize(fyne.NewSize(200, 200))
		fmt.Println(image.Size(), image.File)
		q = append(q, "Throw")
	})
	Print_all := widget.NewButton("Show List", func() {
		fmt.Println(q)
		fmt.Println(q[0])

	})
	//grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 200)), hello, Test2, image, image2)
	top := container.New(layout.NewGridLayout(1), text1, text2, text3)
	grid2 := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 100)), top, Test2, Test3, Print_all, image, sun2)
	w.Resize(fyne.NewSize(500, 500))
	w.SetContent(grid2)
	fmt.Println(v)
	w.ShowAndRun()
}
