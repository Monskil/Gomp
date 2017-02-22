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
	BUTTON_UP      = 0
	BUTTON_DOWN    = 1
	BUTTON_COMMAND = 2
)

type floor_t int

const (
	FLOOR_1 = 0
	FLOOR_2 = 1
	FLOOR_3 = 2
	FLOOR_4 = 3
)

type on_off_t int

const (
	OFF = 0
	ON  = 1
)

type motor_direction_t int

const (
	DIR_DOWN = -1
	DIR_STOP = 0
	DIR_UP   = 1
)

func Set_button_lamp(button button_t, floor floor_t, on_off on_off_t) {
	if on_off == 0 {
		Io_clear_bit(button_channel_matrix[floor][button])
	} else {
		Io_set_bit(button_channel_matrix[floor][button])
	}
}

func Elevator_init() {
	Io_init()

	fmt.Println("Ready to set!")
	// just trying to set some lamps
	Set_button_lamp(BUTTON_UP, FLOOR_1, ON)
	Set_button_lamp(BUTTON_DOWN, FLOOR_2, OFF)
	Set_button_lamp(BUTTON_COMMAND, FLOOR_2, ON)

	// turn all lights off
	// --- can you use int, or do you have to use button_t and floor_t?
	// --- does it work witn int + floor_t?
	//for floor := 0; floor < NUM_FLOORS; floor++ {
	//	for button := 0; button < NUM_BUTTONS; button++ {
	//		Set_button_lamp(button + BUTTON_UP, floor + FLOOR_1, ON)
	//}
	//}
}
