package camera

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/MrRainbow0704/ProgettoIride/internal/config"
)

type VISCACommand string

// Camera commands
const (
	// Reboots the camera, may take some time
	cmdCameraReboot VISCACommand = "81,01,04,00,03,FF"

	// 0x1 = up, 0x2 = down, 0x4 = left, 0x8 = right
	cmdCameraMove VISCACommand = "81,01,04,16,0%c,FF"

	// Accepted values for resolutions are:
	//  - [0x1, 0x3] = 1080p60
	//  - [0x1, 0x4] = 1080p50
	//  - [0x0, 0x6] = 1080p30
	//  - [0x0, 0x8] = 1080p25
	//  - [0x0, 0x9] = 720p60
	//  - [0x0, 0xC] = 720p50
	//  - [0x0, 0xE] = 720p30
	//  - [0x1, 0x1] = 720p25
	cmdCameraResolution VISCACommand = "81,01,04,24,72,0%c,0%c,FF"

	// ON = 0x1, OFF = 0x0
	cmdCameraDualLVDS VISCACommand = "81,01,04,24,74,00,0%c,FF"

	// Accepted values for zoom are:
	//  - from [0x0, 0x0, 0x0, 0x0] to [0x4, 0x0, 0x0, 0x0] for optical zoom
	//  - from [0x4, 0x0, 0x0, 0x1] to [0x7, 0xA, 0xC, 0x0] for digital zoom
	cmdCameraZoom VISCACommand = "81,01,04,47,0%c,0%c,0%c,0%c,FF"

	// ON = 0x2, OFF = 0x3
	cmdCameraMirror VISCACommand = "81,01,04,61,0%c,FF"

	// ON = 0x2, OFF = 0x3
	cmdCameraFlip VISCACommand = "81,01,04,66,0%c,FF"
)

// Camera inquiries
const (
	// Output: [0xA0, 0x50, r1, r2, 0xFF]
	// - r1, r2 = resolution
	inqCameraResolution VISCACommand = "81,09,04,24,72,FF"

	// Output: [0xA0, 0x50, r1, 0xFF]
	// - r1 = 0x2 if mirror is ON, 0x3 if mirror is OFF
	inqCameraMirror VISCACommand = "81,09,04,61,FF"

	// Output: [0xA0, 0x50, r1, 0xFF]
	// - r1 = 0x2 if flip is ON, 0x3 if flip is OFF
	inqCameraFlip VISCACommand = "81,09,04,66,FF"

	// Output: [0xA0, 0x50, r1, r2, r3, r4, 0xff]
	// - r1, r2, r3, r4 = zoom value
	inqCameraZoom VISCACommand = "81,09,04,47,FF"
)

// Camera Interface Board (CIB) commands
const (
	// Resets the CIB, may take some time
	cmdCIBReset VISCACommand = "82,01,0A,00,FF"

	// ON = 0x1, OFF = 0x0
	cmdCIBCrossHair VISCACommand = "82,01,0A,03,0%c,FF"
)

// Cameera Interface Board (CIB) inquiries
const (
	// Output: [0xA0, 0x50, r1, r2, r3, 0xFF]
	//  - r1 = major version
	//  - r2 = minor version
	//  - r3 = sub version
	inqCIBVersion VISCACommand = "82,09,0A,00,FF"

	// Output: [0xA0, 0x50, r1, r2, 0xFF]
	//  - r1 = hardware revision
	//  - r2 = board variant
	inqCIBInfo VISCACommand = "82,09,0A,01,FF"

	// Output: [0xA0, 0x50, r1, r2, 0xFF]
	//  - r1 = status
	//      Bit 0 = Voltage OK
	//      Bit 1 = LVDS PLL clock OK
	//      Bit 2 = Reserved
	//      Bit 3 = Pixel PLL clock OK
	//      Bit 4 = Reserved
	//      Bit 5 = Cam. comms initialized
	//      Bit 6 = Running / OK
	//      Bit 7 = Error state
	//  - r2 =  temperature (+60° offset)
	inqCIBHealth VISCACommand = "82,09,0A,02,FF"

	// Output: [0xA0, 0x50, r1, r2, r3, r4, r5, 0xFF]
	//  - r1 = project code
	//  - r2 = project board
	//  - r3 = board issue
	//  - r4 = build MSB
	//  - r5 = build LSB
	inqCIBHardware VISCACommand = "82,09,0A,04,FF"

	// Expected output: [0xA0, 0x50, r1, 0xFF]
	//  - r1 = error
	//      0x0 = No error
	//      0x1 = FPGA core temperature
	//      0x2 = USB 5V power rail fault
	//      0x3 = Main power rail fault
	//      0x4 = 1V8 power rail fault
	//      0x5 = 3V3 power rail fault
	//      0x6 = 1V1 power rail fault
	//      0x7 = 2V5 power rail fault
	//      0x8 = 5V HDMI DDC fault
	//      0x9 = Camera comms timeout
	//      0xA = Camera video mode/LVDS link width setup fault
	//      0xB = LVDS clock loss of lock
	//      0xC = Reserved
	//      0xD = Pixel clock loss of lock
	//      0xE = USB error
	//      0xF = Firmware type error
	inqCIBError VISCACommand = "82,09,0A,05,FF"
)

func sendVISCACommand(c VISCACommand) error {
	cmd := exec.Command(config.HCPath, "USB3", string(c))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("unable to run command '%s': %w", cmd.String(), err)
	}
	return nil
}

func sendVISCACommandWithOutput(c VISCACommand) ([]byte, error) {
	cmd := exec.Command(config.HCPath, "USB3", string(c))
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("unable to run command '%s': %w", cmd.String(), err)
	}

	_, res, ok := strings.Cut(string(out), "\n")
	if !ok {
		return nil, fmt.Errorf("unable to parse output '%v' of command '%s': no newline found", out, cmd.String())
	}

	bytes := []byte{}
	for b := range strings.SplitSeq(strings.TrimSpace(res[6:len(res)-2]), " ") {
		u, err := strconv.ParseUint(b, 16, 8)
		if err != nil {
			return nil, fmt.Errorf("unable to parse output '%v' of command '%s': %w", out, cmd.String(), err)
		}
		bytes = append(bytes, byte(u))
	}
	return bytes, nil
}
