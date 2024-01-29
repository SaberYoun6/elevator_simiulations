package main

import(
"math"
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
  bottomFloor,finalFloor,currentFloor int
  weight float64
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
func velo(dist0 ,dist1 , time float64) (float64 , error) {
  deltadist := math.Sqrt(math.Pow((dist1-dist0),2))
  velocity  :=  deltadist/time
  if time == 0.0 && dist1 != dist0 {
      return  time , errors.New("Has time stopped and you have broken a the laws of physic or just ripped a wormwhole in the fabric of time and space")
  } else if dist1 == dist0 && time >= 0 {
    return  dist1 , nil
 }
  return velocity, nil
}

func convertingDegreesToRadians(ang float64) (float64,error){
   degreeofPi :=180.0
   deg :=ang/degreeofPi
   rads :=deg*3.1415926
   if rads == math.Inf(-1){
      return rads, errors.New("then you messed up")
   }
   return rads, nil
}

func staticForceFriction(mass , gty , ang float64)  (float64, error) {
   force   := mass * gty
   forceNormal := -force*math.Cos(ang)
   fmt.Println(forceNormal)
   forceNet := force*math.Sin(ang)
   fmt.Println(forceNet)
   finalEnergy := math.Sqrt(math.Abs(forceNet)*2*mass)
   fmt.Println(finalEnergy)
   forFric :=  forceNormal- finalEnergy
   fmt.Println(forFric)
   if force == 0 {
   return  force, errors.New("work cannot be equal to zero")
   }
   return forFric, nil
}
//  func  forFrictCon(forK,forNorm float64) (float64, err) {
//}
/*
func troque(gravity, mass float64 )  (float64, err){
   
}
*/
//func tension(troque, cnt_bal, mat,gravity float64) (float64, err){
//}

/*func strain_cables(ten , gty float64,floor int,) (float64 , float64, err) {
}
*/
/*func strain_ribs(torque, launch_speed, spin_rate, gty float64 ,floor,touch int){

}
*/

func Gravity(g,r float64) float64{
   gravityNearEarthOutside  := math.Pow(r,2)/g
   return  gravityNearEarthOutside
}
func Radius(radii float64) float64 {
  radius:= math.Pow(radii,2)
  return radius
}/*
func tauengine(torque float64 ) float64{
  return torque
}
func taubolts( mass, distance ,angel,gravity float64) float64{
   pg :=  mass*distance* gravity*math.Sin(angel)*math.Cos(90.0-angel)
   return pg
}
func callengineeringfeet()
{
}*/
func force( m, g ,r float64)float64{
  velo := math.Sqrt(r*g/m)
  return  velo
}


// assume that gravity is constant for earth 
// assume that gravity weaks alot at the 80.4672 the furtherest from the orbit
// we are assumeing that we are at  level
func main() {
 floor := 7
 gravity:= -9.81
 topFloor := 10
 botFloor := 0
 ground   := 0.0
 min_dist := 80.4672
 atime    := 900.0
 btime    := 0.0
 ctime    := 120.0
 max_dist := 99.7793
 ext_dist := 115.873
 mass     := 150.00
 ang      :=  15.0
 rad,err  := convertingDegreesToRadians(ang)
 if err != nil {
	 fmt.Println(err)
 }
 fmt.Println(rad)
 vela,err := velo(ground, min_dist,atime)
 if err != nil{
        fmt.Println(err)
 }
  velb,err := velo(ground, ext_dist,btime)
 if err != nil{
        fmt.Println(err)
    }
 velc,err := velo(max_dist,max_dist,ctime)
 if err != nil{
        fmt.Println(err)
 }
 fmt.Printf("The velocity for vela: %.4f\nThe velocity for velb: %.4f\nThe velocity for velc: %.4f\n",vela,velb,velc)
 force,err:=staticForceFriction(mass,gravity,rad)
  if err != nil {
        fmt.Println(err)
 }
 fmt.Printf("the friction force constants is %.3f\n",force)
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
