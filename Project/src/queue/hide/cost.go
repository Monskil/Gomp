// Calculates how much effort it is for the lift to go through with the order.
// With much help from motenfyhn, will miek it less like mortenfyhn l8r if necessary

package queue

import (
        def "config"
        "log"
)

// Function CalculateCost will return the cost. Each sheduled stop and each floor it passes on
// the way towards the selected floor will add 2 to the cost. If the elevator starts between 
// two floor it will add 1. 
//Pseudo code

func CalculateCost(destFloor, destButton, prevFloor, currFloor, currDir int) int {
  q := local.deepCopy()
  q.setOrder(destFloor, def.ButtonInside, orderStatus{true, "", nil})
             
  cost := 0
  floor := prevFloor //moving elevator? VS staning elevator. Why implement floor as prevFloor?
  dir := currDir


  if currFloor == -1 {
      cost ++ //elevator starts between two floors -> adding 1 to cost
  }

  floor, dir = incrementFloor(floor, dor)

// Simulates the elevators journey to the destination floor and calculates the cost until it
// it reaces the destination. Loop from 0 to 10 to make sure its not a infinite loop. 
// Does it consider stops on the way to destination floor? Is that shouldStop?
  
  for n :=0; !(floor == destFLoor && q.shouldStop(floor, dir)) && n<10; n++ {
      if q.shouldStop(floor, dir) {
          cost += 2
          q.setOrder(floor, def.ButtonUp, inactive)
          q.setOrder(floor, def.ButtonDown, inactive)
          q.setOrder(floor, def.ButtonInside, inactive)
      }
  
      dir = q.chooseDirection(floor, dir)
      floor, dir = incrementFloor(floor, dir)
      cost += 2
  }

  return cost
  
}

//returns the floor and direction to the elevator

func incrementFloor(floor, dir int)(int, int) {
  switch dir {
    case def.DirDown:
        floor--
    case def.DirUp:
        floor++
    //case def.DirStop:
    default:
        def.CloseConnectionChan <- true
        def.Restart.Run()
        log.FatalIn(def.ColR, "FAIL", def.ColN)
  }
  
  if floor <= 0 && dir == def.DirDown {
      dir = def = def.DirUp
      floor = 0
  }
  
  if floor >=def.NumFloors-1 && dir == def.DirUp {
       dir = def.DirDown
       floor = de.NumFloors - 1
  }
  
  return floor, dir
  
  }
}
