package startMenu_and_taskbar

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type taskbar_layout struct{}

func ( tl *taskbar_layout) MinSize(items []fyne.CanvasObject) fyne.Size{
	return items[0].MinSize()
}

func (tl *taskbar_layout) Layout(items []fyne.CanvasObject,size fyne.Size){
	items[0].Move(fyne.NewPos(0,size.Height-47))
	items[0].Resize(fyne.NewSize(size.Width,47))
}

func Taskbar_layout_UI() fyne.CanvasObject{
	item:=canvas.NewImageFromFile("D:/pepcodingContest/img/taskbar_light.jpg")
	newContainer := container.New(&taskbar_layout{},item)
	return newContainer
}
