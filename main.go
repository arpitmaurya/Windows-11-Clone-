package main

import (
	news "pepcodingContest/news"
	startMenu "pepcodingContest/startMenu"
	taskbar "pepcodingContest/taskbar"
	taskbarIcon "pepcodingContest/taskbar/taskbar_icon_layoutAndIcon"

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

	myWindow.SetContent( container.NewBorder(nil,nil,nil,nil,wallpapper,taskbar.Taskbar_layout_UI(),
	taskbarIcon.TaskbarIcon_layout_UI(),startMenu.StartMenuLayout_UI(),
	startMenu.StartMenuIconLayout_UI(),news.NewsLayout_UI(),),	

)
	res,_ := fyne.LoadResourceFromPath("D:/pepcodingContest/img/logo.png")
	myWindow.SetIcon(res)

	myWindow.FullScreen()
	myWindow.SetPadded(false)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(1200,680))
	myWindow.ShowAndRun()
}

