package video

import (
	"fmt"
	"image"
	"io"
	"os"
	"os/exec"
	"slices"

	"github.com/MrRainbow0704/ProgettoIride/internal/camera"
)

type LinuxCamera struct {
	id         string
	resolution camera.Resolution
	bufSize    int
	buffer     []byte
	stdout     io.ReadCloser
	cmd        *exec.Cmd
}

var _ Camera = (*LinuxCamera)(nil)

// Returns a list of available cameras by listing the /dev directory
// and filtering for entries that start with "video".
func ListCameras() ([]string, error) {
	dEntry, err := os.ReadDir("/dev")
	if err != nil {
		return nil, fmt.Errorf("failed to list cameras: %w", err)
	}

	var cameras []string
	for _, entry := range dEntry {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if len(name) > 5 && name[:5] == "video" {
			cameras = append(cameras, "/dev/"+name)
		}
	}

	return cameras, nil
}

func NewCamera(id string, res camera.Resolution) *LinuxCamera {
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
	return &LinuxCamera{
		id:         id,
		resolution: res,
		bufSize:    bufSize,
		buffer:     make([]byte, bufSize),
		stdout:     nil,
		cmd:        nil,
	}
}

func (cam *LinuxCamera) StartCapture() error {
	width, height, fps := cam.resolution.Sizes()
	cmd := exec.Command(conf.FFPath,
		"-hide_banner", "-loglevel", "error",
		"-f", "v4l2", "-i", fmt.Sprintf("video=%s", cam.id),
		"-f", "rawvideo",
		"-pix_fmt", "rgba",
		"-s", fmt.Sprintf("%dx%d", width, height),
		"-r", fmt.Sprintf("%d", fps),
		"-",
	)

	var err error
	cam.stdout, err = cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ffmpeg command: %w", err)
	}

	return nil
}

func (cam *LinuxCamera) CaptureBuffer() (*image.RGBA, error) {
	width, height, _ := cam.resolution.Sizes()
	_, err := io.ReadFull(cam.stdout, cam.buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to read frame from stdout: %w", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	copy(img.Pix, cam.buffer)
	return img, nil
}

func (cam *LinuxCamera) StopCapture() error {
	if cam.cmd != nil && cam.cmd.Process != nil {
		if err := cam.cmd.Process.Kill(); err != nil {
			return fmt.Errorf("failed to kill ffmpeg process: %w", err)
		}
	}

	cam.cmd = nil
	cam.stdout = nil
	return nil
}
