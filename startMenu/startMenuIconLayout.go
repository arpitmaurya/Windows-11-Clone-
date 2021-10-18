package startMenuLayout

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/internal/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MenuCheckState int

const (
	CheckOff MenuCheckState = iota
	CheckOn
)

type MenuStateCheck struct {
	widget.BaseWidget
	State MenuCheckState
	normalState string
	openedState string
	// passFunc func()
}

func NewMenuStateCheck(array []string) *MenuStateCheck {
	c := &MenuStateCheck{normalState:array[0],openedState: array[1] }
	c.ExtendBaseWidget(c)
	return c
}

func (c *MenuStateCheck) Tapped(_ *fyne.PointEvent) {
	if c.State == CheckOn {
		c.State = CheckOff
	} else {
		c.State++
	}

	c.Refresh()
}

func (c *MenuStateCheck) CreateRenderer() fyne.WidgetRenderer {
	r := &MenuStateRender{check: c, img: &canvas.Image{}}
	r.updateImage()
	return r
}

type MenuStateRender struct {
	check *MenuStateCheck
	img   *canvas.Image
}

func (t *MenuStateRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (t *MenuStateRender) Destroy() {
}

func (t *MenuStateRender) Layout(_ fyne.Size) {
	t.img.Resize(fyne.NewSize(60,60))
}

func (t *MenuStateRender) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (t *MenuStateRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{t.img}
}

func (t *MenuStateRender) Refresh() {
	t.updateImage()
}

func (t *MenuStateRender) updateImage() {
	defer t.img.Refresh()

	switch t.check.State {
	case CheckOff:
		
		res,err := fyne.LoadResourceFromPath(t.check.normalState)
		if err != nil {
			log.Println("Failed to load indeterminate resource")
			return
		}
		t.img.Resource =theme.NewThemedResource(res)
	default:
		res,err := fyne.LoadResourceFromPath(t.check.openedState)
		if err != nil {
			log.Println("Failed to load indeterminate resource")
			return
		}
		t.img.Resource = theme.NewThemedResource(res)
	}
}

