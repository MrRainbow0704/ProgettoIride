package camera

import (
	"os"
	"path/filepath"
)

var (
	Prog    string
	ProgDir string
)

func init() {
	h, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	ProgDir = filepath.Join(h, ".iride", "bin")
	Prog = filepath.Join(ProgDir, "HarrierControl.exe")
}

func resetCIB() error {
	return sendCommand(cmd_cib_reset)
}

func setDualLVDS() error {
	return sendCommand(cmd_dual_lvds)
}

func setSingleLVDS() error {
	return sendCommand(cmd_single_lvds)
}

func Reboot() error {
	return sendCommand(cmd_reboot)
}
