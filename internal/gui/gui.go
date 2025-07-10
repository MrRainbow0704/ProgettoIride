package gui

import (
	"github.com/MrRainbow0704/ProgettoIride/internal/camera"

	"github.com/lxn/walk"
	d "github.com/lxn/walk/declarative"
)

var mw = new(walk.MainWindow)

func Window() d.MainWindow {
	icon, err := walk.NewIconFromFile("iride.ico")
	if err != nil {
		panic("Inpossibile creare l'immagine")
	}

	videoStream := new(walk.ImageView)

	go videoLoop(videoStream)
	return d.MainWindow{
		AssignTo: &mw,
		Title:    "Iride",
		Icon:     icon,
		MinSize:  d.Size{Width: 600, Height: 400},
		Layout:   d.VBox{},
		MenuItems: []d.MenuItem{
			d.Menu{
				Text: "&Risoluzione",
				Items: []d.MenuItem{
					d.Action{
						Text:        "720p60",
						Enabled:     camera.GetResolution() != camera.R720p60,
						OnTriggered: func() { camera.SetResolution(camera.R720p60) },
					},
					d.Action{
						Text:        "1080p30",
						Enabled:     camera.GetResolution() != camera.R1080p30,
						OnTriggered: func() { camera.SetResolution(camera.R1080p30) },
					},
					d.Action{
						Text:        "1080p60",
						Enabled:     camera.GetResolution() != camera.R1080p60,
						OnTriggered: func() { camera.SetResolution(camera.R1080p60) },
					},
				},
			},
			d.Menu{
				Text: "&Zoom",
				Items: []d.MenuItem{
					d.Action{
						Text:        "Zoom -",
						Shortcut:    d.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyOEMMinus},
						OnTriggered: func() { camera.SetZoom(camera.GetZoom() - 500) },
					},
					d.Action{
						Text:        "Zoom +",
						Shortcut:    d.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyOEMPlus},
						OnTriggered: func() { camera.SetZoom(camera.GetZoom() + 500) },
					},
					d.Action{
						Text:        "Zoom Min",
						Shortcut:    d.Shortcut{Modifiers: walk.ModControl | walk.ModShift, Key: walk.KeyOEMMinus},
						OnTriggered: func() { camera.SetZoomMin() },
					},
					d.Action{
						Text:        "Zoom Max",
						Shortcut:    d.Shortcut{Modifiers: walk.ModControl | walk.ModShift, Key: walk.KeyOEMPlus},
						OnTriggered: func() { camera.SetZoomMax() },
					},
				},
			},
			d.Action{Text: "E&xit", OnTriggered: func() { mw.Close() }},
		},
		Children: []d.Widget{
			d.ImageView{AssignTo: &videoStream, Background: d.SolidColorBrush{Color: walk.RGB(255,0,0)}},
		},
	}
}

func videoLoop(v *walk.ImageView) {
	for {
		// i, err :=
		// if err != nil {
		// 	panic(err)
		// }
		// v.SetImage(i)
		break
	}
}