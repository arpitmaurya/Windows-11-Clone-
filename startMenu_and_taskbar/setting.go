package startMenu_and_taskbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)
var (
	setting_global_width float32
	setting_global_height float32
	setting_itemData fyne.CanvasObject
)
type settingLayout struct{}

func ( nl *settingLayout) MinSize(items []fyne.CanvasObject) fyne.Size{
	return items[0].MinSize()
}

func (nl *settingLayout) Layout(items []fyne.CanvasObject,size fyne.Size){
			items[0].Move(fyne.NewPos(size.Width,size.Height))
		items[0].Resize(fyne.NewSize(size.Width/3.5,180))
		setting_global_width = size.Width
		setting_global_height = size.Height
		setting_itemData = items[0]
}

func Setting_RunUpMove(){
				setting_itemData.Move(fyne.NewPos(setting_global_width/2-190,(setting_global_height/1.5+19)))
}

func Setting_RunUpDown(){
				setting_itemData.Move(fyne.NewPos(setting_global_width,setting_global_height))
}

func SettingLayout_UI() fyne.CanvasObject{
	item:=canvas.NewImageFromFile("D:/pepcodingContest/img/setting_light.png")
	newContainer := container.New(&settingLayout{},item)
	return newContainer
}