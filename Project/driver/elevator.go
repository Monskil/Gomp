package driver

import (
	"fmt"
)

const MOTOR_SPEED int = 2800

var lamp_channel_matrix = [NUM_FLOORS][NUM_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var button_channel_matrix = [NUM_FLOORS][NUM_BUTTONS]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

type button_t int

const (
	BUTTON_UP      = 0
	BUTTON_DOWN    = 1
	BUTTON_COMMAND = 2
)

type motor_direction_t int

const (
	DIR_DOWN = -1
	DIR_STOP = 0
	DIR_UP   = 1
)

func Elevator_init() {
	//MAKING LATER
	io_init()
	Set_button_lamp(0, 0, 0)
	Set_button_lamp(0, 2, 0)
	Set_button_lamp(1, 0, 0)
	Set_button_lamp(1, 1, 0)
	Set_button_lamp(1, 2, 0)
	Set_button_lamp(2, 0, 0)
	Set_button_lamp(2, 1, 0)
	Set_button_lamp(2, 2, 0)
	Set_button_lamp(3, 1, 0)
	Set_button_lamp(3, 2, 0)

}

func Set_button_lamp(button button_t, floor int, on_off int) {
	if on_off == 0 {
		io_clear_bit(button_channel_matrix[floor][button])
	} else {
		io_set_bit(button_channel_matrix[floor][button])
	}
}
