package apps

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Calc(){
	window := fyne.CurrentApp().NewWindow("Calculator")
	
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(270,430))
		img,_ := fyne.LoadResourceFromPath("D:/pepcodingContest/img/calcLogo.png")
	window.SetIcon(img)
	btnTypeLabel := widget.NewEntry()
	btnTypeLabel.SetText(" ")

	ansLabel := canvas.NewText(" 0",color.Black)
	ansLabel.TextSize = 50
	btn0 := widget.NewButton("0",func() {})
	btn1 := widget.NewButton("1",func() {})
	btn2 := widget.NewButton("2",func() {})
	btn3 := widget.NewButton("3",func() {})
	btn4 := widget.NewButton("4",func() {})
	btn5 := widget.NewButton("5",func() {})
	btn6 := widget.NewButton("6",func() {})
	btn7 := widget.NewButton("7",func() {})
	btn8 := widget.NewButton("8",func() {})
	btn9 := widget.NewButton("9",func() {})
	btnDot:=widget.NewButton(".",func() {})
	btnEqual :=widget.NewButton("=",func() {})
	btnEqual.Importance = widget.HighImportance
	btnAdd :=widget.NewButton("+",func() {})
	btnAdd.Importance = widget.HighImportance
	btnSub :=widget.NewButton("-",func() {})
	btnSub.Importance = widget.HighImportance
	btnMul :=widget.NewButton("*",func() {})
	btnMul.Importance = widget.HighImportance
	btnDiv :=widget.NewButton("/",func() {})
	btnDiv.Importance = widget.HighImportance

	btnClear := widget.NewButtonWithIcon("",theme.DeleteIcon(),func() {})
	btnClear.Importance =  widget.HighImportance
	btnSquare := widget.NewButton("x²",func() {})
	btnSquareRoot := widget.NewButton("SQRT",func() {})

	window.SetContent(container.NewVBox(
		layout.NewSpacer(),
		 ansLabel,btnTypeLabel,
			layout.NewSpacer(),
			container.NewGridWrap(
		fyne.NewSize(64.3,55),
		btn7,btn8,btn9,btnDiv,
		btn4,btn5,btn6,btnMul,
		btn1,btn2,btn3,btnSub,
		btnDot,btn0,btnSquare,btnAdd,
		btnSquareRoot,btnClear,btnEqual,
		),layout.NewSpacer(),))
	window.SetPadded(false)
	window.Show()
}