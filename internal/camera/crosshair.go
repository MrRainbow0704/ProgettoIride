package camera

var crosshairStatus bool = false

func CrosshairStatus() bool {
	return crosshairStatus
}

func ToggleCrosshair() error {
	var c cmd
	if crosshairStatus {
		c = cmd_cib_x_hair_off
	} else {
		c = cmd_cib_x_hair_on
	}
	if err := sendCommand(c); err != nil {
		return err
	}
	crosshairStatus = !crosshairStatus
	return nil
}
