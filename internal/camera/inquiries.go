package camera

import (
	"fmt"
	"strconv"
	"strings"
)

type InvalidResponseError struct {
	Response []byte
}

func (e InvalidResponseError) Error() string {
	return fmt.Sprintf("invalid response: %#v", e.Response)
}

func NewInvalidResponseError(response []byte) error {
	return InvalidResponseError{Response: response}
}

type CameraVersion struct {
	Major uint
	Minor uint
	Sub   uint
}

type CameraInfo struct {
	Revision uint
	Variant  uint
}

type CameraHealth struct {
	VoltageOk              bool
	LLVDSPLLClockOk        bool
	PixelPLLClockOk        bool
	CameraCommsInitialized bool
	IsRunning              bool
	Errors                 bool
	Temperature            uint
}

type CameraHardwareInfo struct {
	ProjectCode  uint
	ProjectBoard uint
	BoardIssue   uint
	BuildMSB     uint
	BuildLSB     uint
}

type CameraErrorValue uint

func (c CameraErrorValue) Error() string {
	switch c {
	case 0x0:
		return "No error"
	case 0x1:
		return "FPGA core temperature"
	case 0x2:
		return "USB 5V power rail fault"
	case 0x3:
		return "Main power rail fault"
	case 0x4:
		return "1V8 power rail fault"
	case 0x5:
		return "3V3 power rail fault"
	case 0x6:
		return "1V1 power rail fault"
	case 0x7:
		return "2V5 power rail fault"
	case 0x8:
		return "5V HDMI DDC fault"
	case 0x9:
		return "Camera comms timeout"
	case 0xA:
		return "Camera video mode/LVDS link width setup fault"
	case 0xB:
		return "LVDS clock loss of lock"
	case 0xC:
		return "Reserved"
	case 0xD:
		return "Pixel clock loss of lock"
	case 0xE:
		return "USB error"
	case 0xF:
		return "Firmware type error"
	default:
		return "Unknown error code: " + strings.ToUpper(strconv.FormatUint(uint64(c), 16))
	}
}

func GetVersion() (CameraVersion, error) {
	out, err := sendVISCACommandWithOutput(inqCIBVersion)
	if err != nil {
		return CameraVersion{}, err
	}

	// Expected output: [0xA0, 0x50, r1, r2, r3, 0xFF]
	// r1 = major version
	// r2 = minor version
	// r3 = sub version
	if out[0] != 0x90 || out[1] != 0x50 || out[5] != 0xFF {
		return CameraVersion{}, NewInvalidResponseError(out)
	}
	return CameraVersion{Major: uint(out[2]), Minor: uint(out[3]), Sub: uint(out[4])}, nil
}

func GetInfo() (CameraInfo, error) {
	out, err := sendVISCACommandWithOutput(inqCIBInfo)
	if err != nil {
		return CameraInfo{}, err
	}

	// Expected output: [0xA0, 0x50, r1, r2, 0xFF]
	// r1 = hardware revision
	// r2 = board variant
	if out[0] != 0x90 || out[1] != 0x50 || out[4] != 0xFF {
		return CameraInfo{}, NewInvalidResponseError(out)
	}
	return CameraInfo{Revision: uint(out[2]), Variant: uint(out[3])}, nil
}

func GetHealth() (CameraHealth, error) {
	out, err := sendVISCACommandWithOutput(inqCIBHealth)
	if err != nil {
		return CameraHealth{}, err
	}

	// Expected output: [0xA0, 0x50, r1, r2, 0xFF]
	// r1 = status
	//   Bit 0 = Voltage OK
	//   Bit 1 = LVDS PLL clock OK
	//   Bit 2 = Reserved
	//   Bit 3 = Pixel PLL clock OK
	//   Bit 4 = Reserved
	//   Bit 5 = Cam. comms initialized
	//   Bit 6 = Running / OK
	//   Bit 7 = Error state
	// r2 =  temperature (+60° offset)
	if out[0] != 0x90 || out[1] != 0x50 || out[4] != 0xFF {
		return CameraHealth{}, NewInvalidResponseError(out)
	}
	s := out[2]
	return CameraHealth{
		VoltageOk:              (s & 0b1) != 0,
		LLVDSPLLClockOk:        (s & 0b10) != 0,
		PixelPLLClockOk:        (s & 0b1000) != 0,
		CameraCommsInitialized: (s & 0b100000) != 0,
		IsRunning:              (s & 0b1000000) != 0,
		Errors:                 (s & 0b10000000) != 0,
		Temperature:            uint(out[3]) - 60,
	}, nil
}

func GetHardwareInfo() (CameraHardwareInfo, error) {
	out, err := sendVISCACommandWithOutput(inqCIBHardware)
	if err != nil {
		return CameraHardwareInfo{}, err
	}

	// Expected output: [0xA0, 0x50, r1, r2, r3, r4, r5, 0xFF]
	// r1 = project code
	// r2 = project board
	// r3 = board issue
	// r4 = build MSB
	// r5 = build LSB
	if out[0] != 0x90 || out[1] != 0x50 || out[7] != 0xFF {
		return CameraHardwareInfo{}, NewInvalidResponseError(out)
	}
	return CameraHardwareInfo{
		ProjectCode:  uint(out[2]),
		ProjectBoard: uint(out[3]),
		BoardIssue:   uint(out[4]),
		BuildMSB:     uint(out[5]),
		BuildLSB:     uint(out[6]),
	}, nil
}

func GetErrorValue() (CameraErrorValue, error) {
	out, err := sendVISCACommandWithOutput(inqCIBError)
	if err != nil {
		return 0, err
	}

	// Expected output: [0xA0, 0x50, r1, 0xFF]
	// r1 = error
	if out[0] != 0x90 || out[1] != 0x50 || out[3] != 0xFF {
		return CameraErrorValue(0), NewInvalidResponseError(out)
	}
	return CameraErrorValue(out[2]), nil
}

func getMirror() (bool, error) {
	out, err := sendVISCACommandWithOutput(inqCameraMirror)
	if err != nil {
		return false, err
	}
	
	// Expected output: [0xA0, 0x50, r1, 0xFF]
	// r1 = error
	if out[0] != 0x90 || out[1] != 0x50 || out[3] != 0xFF {
		return false, NewInvalidResponseError(out)
	}
	return out[2] == 0x02, nil
}

func getFlip() (bool, error) {
	out, err := sendVISCACommandWithOutput(inqCameraFlip)
	if err != nil {
		return false, err
	}

	// Expected output: [0xA0, 0x50, r1, 0xFF]
	// r1 = error
	if out[0] != 0x90 || out[1] != 0x50 || out[3] != 0xFF {
		return false, NewInvalidResponseError(out)
	}
	return out[2] == 0x02, nil
}

func getZoom() (uint, error) {
	out, err := sendVISCACommandWithOutput(inqCameraZoom)
	if err != nil {
		return 0, err
	}

	// Expected output: [0xA0, 0x50, r1, r2, r3, r4, 0xFF]
	// r1, r2, r3, r4 = zoom value
	if out[0] != 0x90 || out[1] != 0x50 || out[6] != 0xFF {
		return 0, NewInvalidResponseError(out)
	}

	var z uint
	for i := range 4 {
		z |= uint(out[i+2]&0xF) << ((3 - i) * 4)
	}
	return z, nil
}

func getResolution() (Resolution, error) {
	out, err := sendVISCACommandWithOutput(inqCameraResolution)
	if err != nil {
		return 0, err
	}

	// Expected output: [0xA0, 0x50, r1, r2, 0xFF]
	// r1, r2 = resolution value
	if out[0] != 0x90 || out[1] != 0x50 || out[4] != 0xFF {
		return 0, NewInvalidResponseError(out)
	}

	var z uint
	for i := range 2 {
		z |= uint(out[i+2]&0xF) << ((3 - i) * 4)
	}
	return Resolution(z), nil
}
