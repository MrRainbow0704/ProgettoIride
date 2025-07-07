package camera

var flipStatus bool = false

func FlipStatus() bool {
	return flipStatus
}

func ToggleFlip() error {
	var c cmd
	if flipStatus {
		c = cmd_flip_off
	} else {
		c = cmd_flip_on
	}
	if err := sendCommand(c); err != nil {
		return err
	}
	flipStatus = !flipStatus
	return nil
}
