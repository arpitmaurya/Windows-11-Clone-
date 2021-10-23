package main

import (
	news "pepcodingContest/news"
	startMenu_and_taskbar "pepcodingContest/startMenu_and_taskbar"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())

	myWindow := myApp.NewWindow("Pepcoding Contest Virtual Os - ArpitMaurya")
	wallpapper := canvas.NewImageFromFile("D:/pepcodingContest/img/wallpaper_light.jpg")

	myWindow.SetContent( container.NewBorder(nil,nil,nil,nil,wallpapper,startMenu_and_taskbar.Taskbar_layout_UI(),
	startMenu_and_taskbar.TaskbarIcon_layout_UI(),startMenu_and_taskbar.StartMenuLayout_UI(),
	startMenu_and_taskbar.StartMenuIconLayout_UI(),news.NewsLayout_UI(),startMenu_and_taskbar.SettingLayout_UI(),startMenu_and_taskbar.SettingIconLayout_UI()),)
	res,_ := fyne.LoadResourceFromPath("D:/pepcodingContest/img/logo.png")
	myWindow.SetIcon(res)

	myWindow.FullScreen()
	myWindow.SetPadded(false)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(1200,680))
	myWindow.ShowAndRun()
}

