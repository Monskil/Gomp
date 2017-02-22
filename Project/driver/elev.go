package driver

/*
#include "io.h"
#include "channels.h"
*/

import (
	"fmt"
)

const MOTOR_SPEED int = 2800
const NUM_FLOORS = 4
const NUM_BUTTONS = 3

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

func Set_button_lamp(button button_t, floor floor_t, on_off on_off_t) {
	if on_off == 0 {
		Io_clear_bit(lamp_channel_matrix[floor][button])
	} else {
		Io_set_bit(lamp_channel_matrix[floor][button])
	}
}

func Get_floor_sensor_signal() int {
	if Io_read_bit(SENSOR_FLOOR1) != 0 {
		return 1
	}
	if Io_read_bit(SENSOR_FLOOR2) != 0 {
		return 2
	}
	if Io_read_bit(SENSOR_FLOOR3) != 0 {
		return 3
	}
	if Io_read_bit(SENSOR_FLOOR4) != 0 {
		return 4
	} else {
		return -1
	}
}

func Set_motor_direction(dir motor_direction_t) {
	if dir == DIR_STOP {
		Io_write_analog(MOTOR, 0)
	} else if dir == DIR_UP {
		Io_clear_bit(MOTORDIR)
		Io_write_analog(MOTOR, MOTOR_SPEED)
	} else if dir == DIR_DOWN {
		Io_set_bit(MOTORDIR)
		Io_write_analog(MOTOR, MOTOR_SPEED)
	}
}

func Clear_all_lights() {
	Set_button_lamp(BUTTON_UP, FLOOR_1, OFF)
	Set_button_lamp(BUTTON_UP, FLOOR_2, OFF)
	Set_button_lamp(BUTTON_UP, FLOOR_3, OFF)
	Set_button_lamp(BUTTON_DOWN, FLOOR_2, OFF)
	Set_button_lamp(BUTTON_DOWN, FLOOR_3, OFF)
	Set_button_lamp(BUTTON_DOWN, FLOOR_4, OFF)
	Set_button_lamp(BUTTON_COMMAND, FLOOR_1, OFF)
	Set_button_lamp(BUTTON_COMMAND, FLOOR_2, OFF)
	Set_button_lamp(BUTTON_COMMAND, FLOOR_3, OFF)
	Set_button_lamp(BUTTON_COMMAND, FLOOR_4, OFF)
}

func Elevator_to_first_floor() {
	for Get_floor_sensor_signal() != 1 {
		Set_motor_direction(DIR_DOWN)
	}
	Set_motor_direction(DIR_STOP)
}

func Elevator_init() {
	Io_init()

	fmt.Println("Ready to clear!")
	Clear_all_lights()
	Elevator_to_first_floor()
}
