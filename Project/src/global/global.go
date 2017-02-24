package global

type button_t int

const (
	BUTTON_UP = iota
	BUTTON_DOWN
	BUTTON_COMMAND
)

type floor_t int

const (
	FLOOR_1 = iota
	FLOOR_2
	FLOOR_3
	FLOOR_4
)

type on_off_t int

const (
	OFF = iota
	ON
)

type motor_direction_t int

const (
	DIR_DOWN = -1 << iota
	DIR_STOP
	DIR_UP
)
