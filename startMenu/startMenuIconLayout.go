package startMenuLayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type startMenuIconLayout struct{}


	var (
	startMenu_Icon_global_width float32
	startMenu_Icon_global_height float32
	startMenu_Icon_itemData []fyne.CanvasObject
	startMenu_Icon_global_count int
)


func (tbi_l *startMenuIconLayout) MinSize(items []fyne.CanvasObject) fyne.Size{
	total := fyne.NewSize(0,0)
	for _,obj := range items{
		if !obj.Visible(){
			continue
		}
		total = total.Add(obj.MinSize())
	}
	return total
}

func (tbi_l *startMenuIconLayout) Layout(items []fyne.CanvasObject,size fyne.Size){

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
	
	startMenu_Icon_global_width = size.Width
	startMenu_Icon_global_height = size.Height 
	startMenu_Icon_itemData = items
	startMenu_Icon_global_count = count
	myPos := fyne.NewPos((size.Width/2.55)-float32(count*23),size.Height)
	// myPos := fyne.NewPos((400)-float32(count*23),(400))
			for _,obj := range items{
				if !obj.Visible(){
					continue
				}
				obj.Move(myPos)
				myPos.X = myPos.X+70
				
				obj.Resize(fyne.NewSize(60,60))
			}
}

func StartMenu_Icon_Run_MoveUp(){
	myPos := fyne.NewPos((startMenu_Icon_global_width/2.55)-float32(startMenu_Icon_global_count*23),(startMenu_Icon_global_height/5.5))
					for _,obj := range startMenu_Icon_itemData{
				if !obj.Visible(){
					continue
				}
				obj.Move(myPos)
				myPos.X = myPos.X+70
				
				obj.Resize(fyne.NewSize(60,60))
			}
}
func StartMenu_Icon_Run_MoveDown(){
	myPos := fyne.NewPos((startMenu_Icon_global_width/2.55)-float32(startMenu_Icon_global_count*23),(startMenu_Icon_global_height))
					for _,obj := range startMenu_Icon_itemData{
				if !obj.Visible(){
					continue
				}
				obj.Move(myPos)
				myPos.X = myPos.X+70
				
				obj.Resize(fyne.NewSize(60,60))
			}
}

func StartMenuIconLayout_UI() fyne.CanvasObject{
	 calc := []string{"D:/pepcodingContest/img/calcLogo.png",
		"D:/pepcodingContest/img/calcStarted.png","calc"}

	 codeView := []string{"D:/pepcodingContest/img/codeView.png",
		"D:/pepcodingContest/img/codeViewStarted.png","codeView"}

	 snake := []string{"D:/pepcodingContest/img/snakeLogo.png",
		"D:/pepcodingContest/img/snakeStarted.png","snake"}


	newContainer := container.New(&startMenuIconLayout{},
		NewMenuStateCheck(calc),
		NewMenuStateCheck(codeView),
		NewMenuStateCheck(snake))
	return newContainer
}