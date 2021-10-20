package startMenu_and_taskbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type taskbarIcon_layout struct{}

func (tbi_l *taskbarIcon_layout) MinSize(items []fyne.CanvasObject) fyne.Size{
	total := fyne.NewSize(0,0)
	for _,obj := range items{
		if !obj.Visible(){
			continue
		}
		total = total.Add(obj.MinSize())
	}
	return total
}

func (tbi_l *taskbarIcon_layout) Layout(items []fyne.CanvasObject,size fyne.Size){

	count := -1
	for _,obj := range items{
		if !obj.Visible(){
			continue
		}
		count++
	}
	if count==1{
		count=0
	}
	
	myPos := fyne.NewPos((size.Width/2)-float32(count*23),size.Height-42.5)
			for _,obj := range items{
				if !obj.Visible(){
					continue
				}
				obj.Move(myPos)
				myPos.X = myPos.X+43
				
				obj.Resize(fyne.NewSize(38,38))
			}
}

func TaskbarIcon_layout_UI() fyne.CanvasObject{
	 start := []string{"D:/pepcodingContest/img/logo2.png",
		"D:/pepcodingContest/img/start_logo_started.png","start"}

	 setting := []string{"D:/pepcodingContest/img/settingLogo.png",
		"D:/pepcodingContest/img/settingLogo_started.png","setting"}

	 news := []string{"D:/pepcodingContest/img/newsLogo.png",
		"D:/pepcodingContest/img/newsLogo_Started.png","news"}


	newContainer := container.New(&taskbarIcon_layout{},NewStateCheck(start),NewStateCheck(setting),NewStateCheck(news))
	return newContainer
}