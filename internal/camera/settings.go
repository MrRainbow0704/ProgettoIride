package camera

import (
	"fmt"
	"strings"
)

type resolution uint

func (r resolution) String() string {
	switch r {
	case R720p60:
		return "720p60"
	case R1080p30:
		return "1080p30"
	case R1080p60:
		return "1080p60"
	default:
		return ""
	}
}

const (
	R720p60 resolution = iota
	R1080p30
	R1080p60
)

var (
	crosshairStatus   bool       = false
	flipStatus        bool       = false
	mirrorStatus      bool       = false
	currentResolution resolution = 0
	zoomValue         uint16     = 0
)

func CrosshairStatus() bool {
	return crosshairStatus
}

func ToggleCrosshair() error {
	var c cmd
	if crosshairStatus {
		c = cmd_cib_x_hair_off
	} else {
		c = cmd_cib_x_hair_on
	}
	if err := sendCommand(c); err != nil {
		return err
	}
	crosshairStatus = !crosshairStatus
	return nil
}

func FlipStatus() bool {
	return flipStatus
}

func ToggleFlip() error {
	var c cmd
	if flipStatus {
		c = cmd_flip_off
	} else {
		c = cmd_flip_on
	}
	if err := sendCommand(c); err != nil {
		return err
	}
	flipStatus = !flipStatus
	return nil
}

func MirrorStatus() bool {
	return mirrorStatus
}

func ToggleMirror() error {
	var c cmd
	if mirrorStatus {
		c = cmd_lr_reverse_off
	} else {
		c = cmd_lr_reverse_on
	}
	if err := sendCommand(c); err != nil {
		return err
	}
	mirrorStatus = !mirrorStatus
	return nil
}

func GetResolution() resolution {
	return currentResolution
}

func SetResolution(r resolution) error {
	switch r {
	case R720p60:
		if err := setSingleLVDS(); err != nil {
			return err
		}
		if err := sendCommand(cmd_mode_720p60); err != nil {
			return err
		}
	case R1080p30:
		if err := setSingleLVDS(); err != nil {
			return err
		}
		if err := sendCommand(cmd_mode_1080p30); err != nil {
			return err
		}
	case R1080p60:
		if err := setDualLVDS(); err != nil {
			return err
		}
		if err := sendCommand(cmd_mode_1080p60); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid resolution")
	}
	if err := resetCIB(); err != nil {
		return err
	}
	currentResolution = r
	return nil
}

func GetZoom() uint16 {
	return zoomValue
}

// Sets zoom value, valid range is from 0x0 to 0x7AC0
//   - From 0x0 to 0x4000: Optical zoom
//   - From 0x4001 to 0x7AC0: Digital zoom
func SetZoom(z uint16) error {
	if z >= 0x7AC0 {
		return fmt.Errorf("value too big. Zoom level must be between 0x0 and 0x4000")
	}
	s := strings.Split(fmt.Sprintf("%X", z), "")
	c := cmd(fmt.Sprintf(string(cmd_zoom), s[0], s[1], s[2], s[3]))
	if err := sendCommand(c); err != nil {
		return err
	}
	zoomValue = z
	return nil
}

func SetZoomMin() error {
	return SetZoom(0x0)
}

func SetZoomMax() error {
	return SetZoom(0x4000)
}
