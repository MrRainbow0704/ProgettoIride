package camera

import "fmt"

type resolution uint

func (r resolution) String() string {
	switch r {
	case R720p60:
		return "720p60"
	case R1080p30:
		return "R1080p30"
	case R1080p60:
		return "R1080p60"
	default:
		return ""
	}
}

const (
	R720p60 resolution = iota
	R1080p30
	R1080p60
)

var currentResolution resolution = 0

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
