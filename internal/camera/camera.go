// Il pacchetto camera espone delle funzioni per controllare
// una telecamera ActiveSilicon.
package camera

import (
	"fmt"
	"strconv"
)

type Resolution uint

const (
	RDefault Resolution = 0
	R1080p60 Resolution = 0x13
	R1080p50 Resolution = 0x14
	R1080p30 Resolution = 0x06
	R1080p25 Resolution = 0x08
	R720p60  Resolution = 0x09
	R720p50  Resolution = 0x0C
	R720p30  Resolution = 0x0E
	R720p25  Resolution = 0x11
)

var Resolutions = []Resolution{R1080p60, R1080p50, R1080p30, R1080p25, R720p60, R720p50, R720p30, R720p25}

func (r Resolution) String() string {
	switch r {
	case RDefault:
		return "Default"
	case R1080p60:
		return "1080p60"
	case R1080p50:
		return "1080p50"
	case R1080p30:
		return "1080p30"
	case R1080p25:
		return "1080p25"
	case R720p60:
		return "720p60"
	case R720p50:
		return "720p50"
	case R720p30:
		return "720p30"
	case R720p25:
		return "720p25"
	default:
		return ""
	}
}

// Width, Height, FPS
func (r Resolution) Sizes() (int, int, int) {
	switch r {
	case R1080p60:
		return 1920, 1080, 60
	case R1080p50:
		return 1920, 1080, 50
	case R1080p30:
		return 1920, 1080, 30
	case R1080p25:
		return 1920, 1080, 25
	case R720p60:
		return 1280, 720, 60
	case R720p50:
		return 1280, 720, 50
	case R720p30:
		return 1280, 720, 30
	case R720p25:
		return 1280, 720, 25
	default:
		return 0, 0, 0
	}
}

var (
	crosshairStatus   bool
	flipStatus        bool
	mirrorStatus      bool
	zoomValue         uint
	currentResolution Resolution
)

// Fetches the current values for [flipStatus], [mirrorStatus],
// [zoomValue] and [currentResolution] from the camera.
// It should be called after camera initialization and before
// using any of the other functions in this package.
func FetchCameraStatuses() error {
	var err error
	zoomValue, err = getZoom()
	if err != nil {
		return fmt.Errorf("failed to get zoom value from camera: %w", err)
	}

	flipStatus, err = getFlip()
	if err != nil {
		return fmt.Errorf("failed to get flip value from camera: %w", err)
	}

	mirrorStatus, err = getMirror()
	if err != nil {
		return fmt.Errorf("failed to get mirror value from camera: %w", err)
	}

	currentResolution, err = getResolution()
	if err != nil {
		return fmt.Errorf("failed to get resolution value from camera: %w", err)
	}
	return nil
}

func CrosshairStatus() bool {
	return crosshairStatus
}

func ToggleCrosshair() error {
	v := '1'
	if crosshairStatus {
		v = '0'
	}

	c := VISCACommand(fmt.Sprintf(string(cmdCIBCrossHair), v))
	if err := sendVISCACommand(c); err != nil {
		return err
	}
	crosshairStatus = !crosshairStatus
	return nil
}

func FlipStatus() bool {
	return flipStatus
}

func ToggleFlip() error {
	v := '2'
	if flipStatus {
		v = '3'
	}

	c := VISCACommand(fmt.Sprintf(string(cmdCameraFlip), v))
	if err := sendVISCACommand(c); err != nil {
		return err
	}
	flipStatus = !flipStatus
	return nil
}

func MirrorStatus() bool {
	return mirrorStatus
}

func ToggleMirror() error {
	v := '2'
	if mirrorStatus {
		v = '3'
	}

	c := VISCACommand(fmt.Sprintf(string(cmdCameraMirror), v))
	if err := sendVISCACommand(c); err != nil {
		return err
	}
	mirrorStatus = !mirrorStatus
	return nil
}

func GetResolution() Resolution {
	return currentResolution
}

func SetResolution(r Resolution) error {
	dualLVDS := false
	if r == R1080p60 || r == R1080p50 {
		dualLVDS = true
	}
	if err := setDualLVDS(dualLVDS); err != nil {
		return err
	}

	s := make([]rune, 2)
	for i := range 2 {
		s[1-i] = rune(((uint(r) >> (i * 4)) & 0xF) + '0')
		if s[1-i] > '9' {
			s[1-i] += 7
		}
	}
	c := VISCACommand(fmt.Sprintf(string(cmdCameraResolution), s[0], s[1]))
	if err := sendVISCACommand(c); err != nil {
		return err
	}

	if err := resetCIB(); err != nil {
		return err
	}
	currentResolution = r
	return nil
}

func GetZoom() uint {
	return zoomValue
}

// Sets zoom value, valid range is from 0x0 to 0x7AC0
//   - From 0x0 to 0x4000: Optical zoom
//   - From 0x4001 to 0x7AC0: Digital zoom
func SetZoom(z uint) error {
	if z > 0x7AC0 {
		return fmt.Errorf("value too big. Zoom level must be between 0x0 and 0x7AC0. Is 0x%X", z)
	}

	s := make([]rune, 4)
	for i := range 4 {
		s[3-i] = rune(((z >> (i * 4)) & 0xF) + '0')
		if s[3-i] > '9' {
			s[3-i] += 7
		}
	}
	c := VISCACommand(fmt.Sprintf(string(cmdCameraZoom), s[0], s[1], s[2], s[3]))
	if err := sendVISCACommand(c); err != nil {
		return err
	}
	zoomValue = z
	return nil
}

func IsZoomDigital(z uint) bool {
	return 0x4000 < z && z <= 0x7AC0
}

func IsZoomOptical(z uint) bool {
	return z <= 0x4000
}

// Sets zoom to the minimum value (0x0)
func SetZoomMin() error {
	return SetZoom(0x0)
}

// Sets zoom to the maximum value (0x4000)
func SetZoomMax() error {
	return SetZoom(0x4000)
}

func resetCIB() error {
	return sendVISCACommand(cmdCIBReset)
}

func setDualLVDS(s bool) error {
	v := '0'
	if s {
		v = '1'
	}

	c := VISCACommand(fmt.Sprintf(string(cmdCameraDualLVDS), v))
	return sendVISCACommand(c)
}

func Reboot() error {
	return sendVISCACommand(cmdCameraReboot)
}

type Direction uint

const (
	Up Direction = 1 << iota
	Down
	Left
	Right
)

func MoveUp() error { return move(Up) }

func MoveDown() error { return move(Down) }

func MoveLeft() error { return move(Left) }

func MoveRight() error { return move(Right) }

func move(d Direction) error {
	s := []rune(strconv.FormatUint(uint64(d), 16))
	c := VISCACommand(fmt.Sprintf(string(cmdCameraMove), s[0]))
	return sendVISCACommand(c)
}
