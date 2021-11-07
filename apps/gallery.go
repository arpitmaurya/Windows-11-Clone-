package apps

import (
	"fmt"
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Gallery() {

	window := fyne.CurrentApp().NewWindow("Gallery")
	r, _ := fyne.LoadResourceFromPath("D:\\Windows_11_clone\\img\\gallery.png")
	window.SetIcon(r)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(500, 500))

	location := "D:\\Windows_11_clone\\gallery_img"

	files, _ := ioutil.ReadDir(location)

	tabs := container.NewAppTabs()
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpg" {
				image := canvas.NewImageFromFile(location + "\\" + file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	window.SetContent(tabs)
	window.Show()
}
