package apps

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Store(){
	window := fyne.CurrentApp().NewWindow("Microsoft Store")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(1100,620))
		img,_ := fyne.LoadResourceFromPath("D:/pepcodingContest/img/store_light.png")
	window.SetIcon(img)
	window.SetPadded(false)
	imgPage := canvas.NewImageFromFile("D:/pepcodingContest/img/storePage.jpg")
	imgPage.Resize(fyne.NewSize(1100,620))
	window.SetContent(container.NewWithoutLayout(imgPage))

	window.Show()
}