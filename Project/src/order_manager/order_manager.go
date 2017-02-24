package order_manager

import (
	"driver"
	"fmt"
)

func Button_two() {
	for {
		if driver.Get_button_signal(BUTTON_UP, FLOOR_2) {
			fmt.Println("Hello, you pressed button up floor 2 hehe.")
		}
	}
}
