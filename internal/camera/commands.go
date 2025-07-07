package camera

import (
	"os/exec"
	"strconv"
	"strings"
)

type cmd string

const (
	cmd_reboot         cmd = "81,01,04,00,03,FF"
	cmd_move_up        cmd = "81,01,04,16,01,FF"
	cmd_move_down      cmd = "81,01,04,16,02,FF"
	cmd_move_left      cmd = "81,01,04,16,04,FF"
	cmd_move_right     cmd = "81,01,04,16,08,FF"
	cmd_mode_1080p60   cmd = "81,01,04,24,72,01,03,FF"
	cmd_mode_1080p30   cmd = "81,01,04,24,72,00,06,FF"
	cmd_mode_720p60    cmd = "81,01,04,24,72,00,09,FF"
	cmd_dual_lvds      cmd = "81,01,04,24,74,00,01,FF"
	cmd_single_lvds    cmd = "81,01,04,24,74,00,00,FF"
	cmd_zoom           cmd = "81,01,04,47,0%s,0%s,0%s,0%s,FF"
	colour_hue         cmd = "81,01,04,4F,00,00,00,00,FF"
	cmd_lr_reverse_on  cmd = "81,01,04,61,02,FF"
	cmd_lr_reverse_off cmd = "81,01,04,61,03,FF"
	cmd_flip_on        cmd = "81,01,04,66,02,FF"
	cmd_flip_off       cmd = "81,01,04,66,03,FF"
	cmd_text_on        cmd = "81,01,04,74,2F,FF"
	cmd_text_off       cmd = "81,01,04,74,3F,FF"
	cmd_cib_reset      cmd = "82,01,0A,00,FF"
	cmd_cib_x_hair_on  cmd = "82,01,0A,03,01,FF"
	cmd_cib_x_hair_off cmd = "82,01,0A,03,00,FF"
	query_version      cmd = "82,09,0A,00,FF"
	query_info         cmd = "82,09,0A,01,FF"
	query_health       cmd = "82,09,0A,02,FF"
	query_error        cmd = "82,09,0A,05,FF"
)

func sendCommand(c cmd) error {
	if err := exec.Command(Prog, "USB3", string(c)).Run(); err != nil {
		return err
	}
	return nil
}

func sendCommandWithOutput(c cmd) ([]uint8, error) {
	out, err := exec.Command(Prog, "USB3", string(c)).Output()
	if err != nil {
		return nil, err
	}
	bytes := []uint8{}
	for b := range strings.SplitSeq(strings.TrimSpace(string(out)), " ") {
		byte, err := strconv.ParseUint(b, 16, 8)
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, uint8(byte))
	}
	return bytes, nil
}
