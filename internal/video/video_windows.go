package video

import (
	"errors"
	"fmt"
	"image"
	"io"
	"os/exec"
	"regexp"
	"slices"

	"github.com/MrRainbow0704/ProgettoIride/internal/camera"
)

type WindowsCamera struct {
	id         string
	resolution camera.Resolution
	bufSize    int
	buffer     []byte
	stdout     io.ReadCloser
	cmd        *exec.Cmd
}

var _ Camera = (*WindowsCamera)(nil)

// Returns a list of available cameras using ffmpeg and the dshow input device.
// It runs the command `ffmpeg -list_devices true -f dshow -i ""`.
func ListCameras() ([]string, error) {
	cmd := exec.Command(conf.FFPath, "-list_devices", "true", "-f", "dshow", "-i", `""`)

	var stdout []byte
	var err error
	if stdout, err = cmd.CombinedOutput(); err != nil && errors.Is(err, &exec.ExitError{}) {
		return nil, fmt.Errorf("failed to run ffmpeg command '%s': %w", cmd.String(), err)
	}

	var cameras []string
	pattern := regexp.MustCompile(`(?m)^\[dshow @ [0-9a-f]{16}\] "(.*)" `)
	for _, matches := range pattern.FindAllSubmatch(stdout, -1) {
		cameras = append(cameras, string(matches[1]))
	}
	return cameras, nil
}

func NewCamera(id string, res camera.Resolution) *WindowsCamera {
	var cams []string
	var err error
	if cams, err = ListCameras(); err != nil {
		return nil
	}
	if !slices.Contains(cams, id) {
		return nil
	}
	width, height, _ := res.Sizes()
	bufSize := width * height * 4
	return &WindowsCamera{
		id:         id,
		resolution: res,
		bufSize:    bufSize,
		buffer:     make([]byte, bufSize),
		stdout:     nil,
		cmd:        nil,
	}
}

func (cam *WindowsCamera) StartCapture() error {
	width, height, fps := cam.resolution.Sizes()
	cam.cmd = exec.Command(conf.FFPath,
		"-hide_banner", "-loglevel", "error",
		"-f", "dshow", "-i", fmt.Sprintf("video=%s", cam.id),
		"-f", "rawvideo",
		"-pix_fmt", "rgba",
		"-s", fmt.Sprintf("%dx%d", width, height),
		"-r", fmt.Sprintf("%d", fps),
		"-",
	)

	var err error
	cam.stdout, err = cam.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	if err := cam.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ffmpeg command: %w", err)
	}

	return nil
}

func (cam *WindowsCamera) CaptureBuffer() (*image.RGBA, error) {
	width, height, _ := cam.resolution.Sizes()
	_, err := io.ReadFull(cam.stdout, cam.buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to read frame from stdout: %w", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	copy(img.Pix, cam.buffer)
	return img, nil
}

func (cam *WindowsCamera) StopCapture() error {
	if cam.cmd != nil && cam.cmd.Process != nil {
		if err := cam.cmd.Process.Kill(); err != nil {
			return fmt.Errorf("failed to kill ffmpeg process: %w", err)
		}
	}

	cam.cmd = nil
	cam.stdout = nil
	return nil
}
