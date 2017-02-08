package main

import (
 	"fmt"
	"log"
	"net"
	"time"
)


func check_for_error(err error){
	if err != nil{
		log.Fatal(err)
	}
}
