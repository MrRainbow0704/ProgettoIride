package camera

var mirrorStatus bool = false

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
