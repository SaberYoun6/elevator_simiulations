package main

import(
//"math"
"fmt"
"errors"
//"bufio"
)
type FloorStruct struct {
     name string
     length , height , width float64
}
type SubStruct struct{
	name string
	horizontal , veritcal float64
}

type Design struct{
  name string
  bottomFloor,finalFloor int
}

func currentFloor(floor, botFloor, topFloor int) (int , error){
    // define what floor is being is being called how it called
    // how high up it goes
    if floor <  botFloor || floor > topFloor {
    return 0, errors.New("cannot paritucally call any floor that is lower then the floor that is already stated")
   } else if  floor > botFloor && floor < topFloor {
     return  floor, nil
   } else if floor == botFloor || floor == topFloor {
     return  floor, nil
   } else{
     return floor , errors.New(" I cannotn think of a current error that could come up with this but I'm sure there is.")
  }
}
func goingUp(currentFloor, floorUpward, topFloor int) (int , error){
     currentfloor :=  currentFloor + floorUpward
     if currentfloor > topFloor{
        return topFloor, errors.New("You have exceed the maxium amount of floors")
     } else {
        return currentfloor, nil
     }
}
func goingDown(currentFloor, floorDownard , botFloor int)( int , error) {
    currentfloor := currentFloor - floorDownard
    if currentfloor < botFloor{
    return botFloor, errors.New("You cannot exceed the lowest floor")
    } else {
      return  currentfloor, nil
    }
}
/*
func gravity(g,r float64) float64{
   e  := math.Pow(r,2)/g
   return  e
}
func radius(radii float64) float64 {
  e:= math.Pow(radii,2)
  return e
}
func tauengine(torque float64 ) float64{
  return torque
}
func taubolts( mass, distance ,angel,gravity float64) float64{
   pg :=  mass*distance* gravity*math.Sin(angel)*math.Cos(90.0-angel)
   return pg
}

func callengineeringfeet()
func force( m, g ,r float64)float64{
  velo := math.Sqrt(r*g/m)
  return  velo
}
*/
func main() {
 floor := 7
 topFloor := 10
 botFloor := 0
 current,err:=currentFloor(floor,botFloor,topFloor)
 if err != nil {
        fmt.Println(err)
 }
 goingup,err := goingUp(current,2,topFloor)
 if err != nil {
       fmt.Println(err)
 }
 goingdown,err := goingDown(goingup,6,botFloor)
 if err != nil {
      fmt.Println(err)
 }
 fmt.Println(goingdown)
 }
