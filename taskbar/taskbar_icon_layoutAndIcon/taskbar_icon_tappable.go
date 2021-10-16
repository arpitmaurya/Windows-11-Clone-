package taskbarIcon

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/internal/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CheckState int

const (
	CheckOff CheckState = iota
	CheckOn
)

type StateCheck struct {
	widget.BaseWidget
	State CheckState
	normalState string
	openedState string
	// passFunc func()
}

func NewStateCheck(array []string) *StateCheck {
	c := &StateCheck{normalState:array[0],openedState: array[1] }
	c.ExtendBaseWidget(c)
	return c
}

func (c *StateCheck) Tapped(_ *fyne.PointEvent) {
	if c.State == CheckOn {
		c.State = CheckOff
	} else {
		c.State++
	}

	c.Refresh()
}

func (c *StateCheck) CreateRenderer() fyne.WidgetRenderer {
	r := &StateRender{check: c, img: &canvas.Image{}}
	r.updateImage()
	return r
}

type StateRender struct {
	check *StateCheck
	img   *canvas.Image
}

func (t *StateRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (t *StateRender) Destroy() {
}

func (t *StateRender) Layout(_ fyne.Size) {
	t.img.Resize(fyne.NewSize(38,38))
}

func (t *StateRender) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (t *StateRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{t.img}
}

func (t *StateRender) Refresh() {
	t.updateImage()
}

func (t *StateRender) updateImage() {
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

