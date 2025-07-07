package camera

import (
	"fmt"
	"strings"
)

var zoomValue uint16 = 0

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
