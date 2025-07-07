package camera

func MoveUp() error {
	return sendCommand(cmd_move_up)
}

func MoveDown() error {
	return sendCommand(cmd_move_down)
}

func MoveLeft() error {
	return sendCommand(cmd_move_left)
}

func MoveRight() error {
	return sendCommand(cmd_move_right)
}
