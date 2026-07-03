// Il pacchetto video espone le funzioni per la gestione
// della cattura video da una webcam
package video

import (
	"fmt"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/MrRainbow0704/ProgettoIride/internal/config"
	"github.com/MrRainbow0704/ProgettoIride/internal/log"
)

var conf = config.Get()
var stopChan chan struct{} = make(chan struct{})
var isBackgroundVideoLoopRunning = false

type Camera interface {
	StartCapture() error
	CaptureBuffer() (*image.RGBA, error)
	StopCapture() error
}

func backgroundVideoLoop(cam Camera, ev *canvas.Image, stop <-chan struct{}) {
	defer cam.StopCapture()

	for {
		select {
		case <-stop:
			log.Info("stopping camera capture...")
			return
		default:
			frame, err := cam.CaptureBuffer()
			if err != nil {
				log.Errorf("failed to capture frame: %v", err)
				continue
			}

			if ev == nil || ev.Image == nil {
				log.Info("embedded video image is nil, skipping frame update")
				continue
			}

			fyne.Do(func() {
				ev.Image = frame
				ev.Refresh()
			})
		}
	}
}

// Starts a gorutine that constatly fetches frames from the given camera
// and updates the given image with the frame.
// Can be stopped by calling [StopBackgroundVideoLoop].
func StartBackgroundVideoLoop(cam Camera, ev *canvas.Image) error {
	if isBackgroundVideoLoopRunning {
		return fmt.Errorf("there can only be one BackgroundVideoLoop running at a time")
	}

	err := cam.StartCapture()
	if err != nil {
		return err
	}

	go backgroundVideoLoop(cam, ev, stopChan)
	isBackgroundVideoLoopRunning = true

	return nil
}

// Stops the gorutine started by [StartBackgroundVideoLoop].
func StopBackgroundVideoLoop() {
	stopChan <- struct{}{}
	isBackgroundVideoLoopRunning = false
}
