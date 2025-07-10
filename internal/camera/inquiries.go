package camera

import "fmt"

type CameraVersion struct {
	Major uint8
	Minor uint8
	Sub   uint8
}

type CameraInfo struct {
	Revision uint8
	Variant  uint8
}

type CameraHealth struct {
	VoltageOk              bool
	LLVDSPLLClockOk        bool
	PixelPLLClockOk        bool
	CameraCommsInitialized bool
	IsRunning              bool
	Errors                 bool
	Temperature            uint8
}

type CameraHardwareInfo struct {
	ProjectCode  uint8
	ProjectBoard uint8
	BoardIssue   uint8
	BuildMSB     uint8
	BuildLSB     uint8
}

type CameraError uint8

func (c CameraError) Error() string {
	switch c {
	case 0x00:
		return "No error"
	case 0x01:
		return "FPGA core temperature"
	case 0x02:
		return "USB 5V power rail fault"
	case 0x03:
		return "Main power rail fault"
	case 0x04:
		return "1V8 power rail fault"
	case 0x05:
		return "3V3 power rail fault"
	case 0x06:
		return "1V1 power rail fault"
	case 0x07:
		return "2V5 power rail fault"
	case 0x08:
		return "5V HDMI DDC fault"
	case 0x09:
		return "Camera comms timeout"
	case 0x0A:
		return "Camera video mode/LVDS link width setup fault"
	case 0x0B:
		return "LVDS clock loss of lock"
	case 0x0C:
		return "Reserved"
	case 0x0D:
		return "Pixel clock loss of lock"
	case 0x0E:
		return "USB error"
	case 0x0F:
		return "Firmware type error"
	default:
		return ""
	}
}

func GetVersion() (CameraVersion, error) {
	out, err := sendCommandWithOutput(cmd_query_version)
	if err != nil {
		return CameraVersion{}, err
	}
	// Expected output: [0xA0, 0x50, r1, r2, r3, 0xFF]
	// r1 = major version
	// r2 = minor version
	// r3 = sub version
	if out[0] != 0xA0 || out[1] != 0x50 || out[5] != 0xFF {
		return CameraVersion{}, fmt.Errorf("la risposta ricevuta, %+v, non è valida", out)
	}
	return CameraVersion{Major: out[2], Minor: out[3], Sub: out[4]}, nil
}

func GetInfo() (CameraInfo, error) {
	out, err := sendCommandWithOutput(cmd_query_info)
	if err != nil {
		return CameraInfo{}, err
	}
	// Expected output: [0xA0, 0x50, r1, r2, 0xFF]
	// r1 = hardware revision
	// r2 = board variant
	if out[0] != 0xA0 || out[1] != 0x50 || out[4] != 0xFF {
		return CameraInfo{}, fmt.Errorf("la risposta ricevuta, %+v, non è valida", out)
	}
	return CameraInfo{Revision: out[2], Variant: out[3]}, nil
}

func GetHealth() (CameraHealth, error) {
	out, err := sendCommandWithOutput(cmd_query_health)
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
	if out[0] != 0xA0 || out[1] != 0x50 || out[4] != 0xFF {
		return CameraHealth{}, fmt.Errorf("la risposta ricevuta, %+v, non è valida", out)
	}
	s := out[2]
	return CameraHealth{
		VoltageOk:              (s & 0b1) != 0,
		LLVDSPLLClockOk:        (s & 0b10) != 0,
		PixelPLLClockOk:        (s & 0b1000) != 0,
		CameraCommsInitialized: (s & 0b100000) != 0,
		IsRunning:              (s & 0b1000000) != 0,
		Errors:                 (s & 0b10000000) != 0,
		Temperature:            out[3] - 60,
	}, nil
}

func GetHardwareInfo() (CameraHardwareInfo, error) {
	out, err := sendCommandWithOutput(cmd_query_health)
	if err != nil {
		return CameraHardwareInfo{}, err
	}
	// Expected output: [0xA0, 0x50, r1, r2, r3, r4, r5, 0xFF]
	// r1 = project code
	// r2 = project board
	// r3 = board issue
	// r4 = build MSB
	// r5 = build LSB
	if out[0] != 0xA0 || out[1] != 0x50 || out[7] != 0xFF {
		return CameraHardwareInfo{}, fmt.Errorf("la risposta ricevuta, %+v, non è valida", out)
	}
	return CameraHardwareInfo{
		ProjectCode:  out[2],
		ProjectBoard: out[3],
		BoardIssue:   out[4],
		BuildMSB:     out[5],
		BuildLSB:     out[6],
	}, nil
}

func GetError() (CameraError, error) {
	out, err := sendCommandWithOutput(cmd_query_info)
	if err != nil {
		return 0, err
	}
	// Expected output: [0xA0, 0x50, r1, 0xFF]
	// r1 = error
	if out[0] != 0xA0 || out[1] != 0x50 || out[3] != 0xFF {
		return 0, fmt.Errorf("la risposta ricevuta, %+v, non è valida", out)
	}
	return CameraError(out[2]), nil
}
