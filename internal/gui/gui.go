// Il pacchetto gui espone le funzioni per la gestione
// dell'interfaccia grafica dell'applicazione
package gui

import (
	"embed"
	"fmt"
	"image"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/MrRainbow0704/ProgettoIride/internal/camera"
	"github.com/MrRainbow0704/ProgettoIride/internal/config"
	"github.com/MrRainbow0704/ProgettoIride/internal/gui/assets"
	"github.com/MrRainbow0704/ProgettoIride/internal/log"
	"github.com/MrRainbow0704/ProgettoIride/internal/version"
	"github.com/MrRainbow0704/ProgettoIride/internal/video"
)

//go:embed translations/*.json
var translationsFS embed.FS

var EmbeddedVideo = canvas.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 1000, 1000)))
var conf = config.Get()

func GUI() fyne.Window {
	a := app.NewWithID("io.github.MrRainbow0704.iride")
	a.SetIcon(assets.ResourceIrideIco)
	w := a.NewWindow(lang.L("iride"))
	w.Resize(fyne.NewSize(800, 600))

	if err := lang.AddTranslationsFS(translationsFS, "translations"); err != nil {
		log.Errorf("failed to load translations: %v", err)
	}

	// Toolbar
	// Bottoni per il toggle di mirror, flip e crosshair
	mirrorToggle := widget.NewCheck(lang.L("mirror"), toggleMirror)
	mirrorToggle.Checked = camera.MirrorStatus()

	flipToggle := widget.NewCheck(lang.L("flip"), toggleFlip)
	flipToggle.Checked = camera.FlipStatus()

	crosshairToggle := widget.NewCheck(lang.L("crosshair"), toggleCrosshair)
	crosshairToggle.Checked = camera.CrosshairStatus()

	// Bottoni per lo zoom
	zoomAmount := uint(0x800)
	zoomLabel := widget.NewLabel(fmt.Sprintf("%d%%", camera.GetZoom()/(100*0x4000)))
	zoomOutBtn := widget.NewButtonWithIcon("", theme.ZoomOutIcon(), zoomOut(zoomLabel, zoomAmount))
	zoomInBtn := widget.NewButtonWithIcon("", theme.ZoomInIcon(), zoomIn(zoomLabel, zoomAmount))
	if camera.GetZoom() > 0x4000 {
		zoomLabel.Importance = widget.WarningImportance
	}

	// Bottone per aprire la finestra di configurazione
	configBtn := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() { openConfigWindow(a) })

	// Bottone per riavviare la camera
	rebootBtn := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), rebootCamera)
	rebootBtn.Importance = widget.DangerImportance

	toolbar := container.NewHBox(
		mirrorToggle,
		flipToggle,
		crosshairToggle,
		zoomOutBtn,
		zoomLabel,
		zoomInBtn,
		rebootBtn,
		layout.NewSpacer(),
		configBtn,
	)

	EmbeddedVideo.FillMode = canvas.ImageFillContain

	mainLayout := container.NewBorder(toolbar, nil, nil, nil, EmbeddedVideo)

	// Shortcuts
	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch ke.Name {
		case fyne.KeyW, fyne.KeyUp:
			moveUp()
		case fyne.KeyA, fyne.KeyLeft:
			moveLeft()
		case fyne.KeyS, fyne.KeyDown:
			moveDown()
		case fyne.KeyD, fyne.KeyRight:
			moveRight()
		case fyne.KeyPlus:
			zoomIn(zoomLabel, zoomAmount)()
		case fyne.KeyMinus:
			zoomOut(zoomLabel, zoomAmount)()
		case fyne.KeyX:
			toggleCrosshair(!crosshairToggle.Checked)
			crosshairToggle.Checked = !crosshairToggle.Checked
		case fyne.KeyM:
			toggleMirror(!mirrorToggle.Checked)
			mirrorToggle.Checked = !mirrorToggle.Checked
		case fyne.KeyF:
			toggleFlip(!flipToggle.Checked)
			flipToggle.Checked = !flipToggle.Checked
		}
	})

	w.SetContent(mainLayout)
	return w
}

func openConfigWindow(a fyne.App) {
	w := a.NewWindow(lang.L("iride") + " - " + lang.L("settings"))
	w.Resize(fyne.NewSize(640, 320))
	w.SetFixedSize(true)

	// Form per la configurazione della camera
	cams, err := video.ListCameras()
	if err != nil {
		log.Errorf("Failed to list cameras: %v", err)
		cams = []string{}
	}
	formEntryDeviceID := widget.NewSelect(cams, func(selected string) {})
	formEntryDeviceID.SetSelected(conf.DeviceID)
	formEntryDeviceID.PlaceHolder = lang.L("placeholder_camera")

	formEntryFFPath := widget.NewEntry()
	formEntryFFPath.SetPlaceHolder(lang.L("placeholder_FFPath"))
	formEntryFFPath.SetText(conf.FFPath)
	formEntryFFPath.AlwaysShowValidationError = true
	formEntryFFPath.Validator = func(s string) error {
		if s == "" {
			return fmt.Errorf("%s", lang.L("errror_empty_ffmpeg"))
		}
		if _, err := exec.LookPath(s); err != nil {
			return fmt.Errorf("%s", lang.L("errror_invalid_ffmpeg"))
		}
		return nil
	}

	resolutions := make([]string, len(camera.Resolutions))
	resolutionsMap := make(map[string]camera.Resolution, len(camera.Resolutions))
	for i, res := range camera.Resolutions {
		resolutions[i] = res.String()
		resolutionsMap[res.String()] = res
	}
	formSelectResolution := widget.NewSelect(resolutions, func(selected string) {})
	formSelectResolution.SetSelected(camera.GetResolution().String())
	formSelectResolution.PlaceHolder = lang.L("placeholder_resolution")

	form := widget.NewForm(
		widget.NewFormItem(lang.L("form_camera_label"), formEntryDeviceID),
		widget.NewFormItem(lang.L("form_ffmpeg_label"), formEntryFFPath),
		widget.NewFormItem(lang.L("form_resolution_label"), formSelectResolution),
	)
	form.OnSubmit = func() {
		// salva la configurazione e riavvia il loop video con
		// la nuova configurazione
		conf.DeviceID = formEntryDeviceID.Selected
		conf.FFPath = formEntryFFPath.Text
		camera.SetResolution(resolutionsMap[formSelectResolution.Selected])
		if err := config.SaveConfig(); err != nil {
			log.Errorf("Failed to save configuration: %v", err)
		} else {
			log.Infof("Saved new configuration: %+v", conf)
			video.StopBackgroundVideoLoop()
			cam := video.NewCamera(conf.DeviceID, camera.GetResolution())
			if cam == nil {
				log.Errorf("Failed to initialize camera with device id: %s", conf.DeviceID)
				return
			}
			if err := camera.FetchCameraStatuses(); err != nil {
				log.Errorf("failed to fetch camera statuses: %v", err)
				return
			}
			if err := video.StartBackgroundVideoLoop(cam, EmbeddedVideo); err != nil {
				log.Errorf("Failed to start background video loop: %v", err)
			}
		}
		w.Close()
	}

	var versionLabel *widget.Label
	if version.IsDev() {
		versionLabel = widget.NewLabel(fmt.Sprintf("%s: %s (%s)", lang.L("version"), version.Get(), lang.L("dev_build")))
	} else {
		versionLabel = widget.NewLabel(fmt.Sprintf("%s: %s", lang.L("version"), version.Get()))
	}

	content := container.NewVBox(form, layout.NewSpacer(), container.NewCenter(versionLabel))

	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeyEscape {
			w.Close()
		}
	})

	w.SetContent(content)

	w.Show()
	w.RequestFocus()
}

func toggleMirror(state bool) {
	if err := camera.ToggleMirror(); err != nil {
		log.Errorf("unable to toggle mirror: %v", err)
	}
}

func toggleFlip(state bool) {
	if err := camera.ToggleFlip(); err != nil {
		log.Errorf("unable to toggle flip: %v", err)
	}
}

func toggleCrosshair(state bool) {
	if err := camera.ToggleCrosshair(); err != nil {
		log.Errorf("unable to toggle crosshair: %v", err)
	}
}

func zoomIn(l *widget.Label, z uint) func() {
	return func() {
		if err := camera.SetZoom(camera.GetZoom() + z); err != nil {
			log.Errorf("unable to zoom in: %v", err)
		} else {
			l.SetText(fmt.Sprintf("%d%%", camera.GetZoom()/(100*0x4000)))
			if camera.IsZoomDigital(camera.GetZoom()) {
				l.Importance = widget.WarningImportance
			} else {
				l.Importance = widget.MediumImportance
			}
		}
	}
}

func zoomOut(l *widget.Label, z uint) func() {
	return func() {
		if err := camera.SetZoom(camera.GetZoom() - z); err != nil {
			log.Errorf("unable to zoom out: %v", err)
		} else {
			l.SetText(fmt.Sprintf("%d%%", camera.GetZoom()/(100*0x4000)))
			if camera.IsZoomDigital(camera.GetZoom()) {
				l.Importance = widget.WarningImportance
			} else {
				l.Importance = widget.MediumImportance
			}
		}
	}
}

func moveUp() {
	if err := camera.MoveUp(); err != nil {
		log.Errorf("unable to move up: %v", err)
	}
}

func moveDown() {
	if err := camera.MoveDown(); err != nil {
		log.Errorf("unable to move down: %v", err)
	}
}

func moveLeft() {
	if err := camera.MoveLeft(); err != nil {
		log.Errorf("unable to move left: %v", err)
	}
}

func moveRight() {
	if err := camera.MoveRight(); err != nil {
		log.Errorf("unable to move right: %v", err)
	}
}

func rebootCamera() {
	if err := camera.Reboot(); err != nil {
		log.Errorf("unable to reboot camera: %v", err)
	}
}
