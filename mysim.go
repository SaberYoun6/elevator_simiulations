package main

import "math"
import "fmt"

type ele struct{
  name string
  bottomFloor,finalFloor float64
}

func isGoingUp(x,y float64)  float64{
   if x >= y{
     return x-y
   }
   return y-x
}
func isGoingDown(x,y float64)  float64{
   if x<=y {
     return math.Abs(x-y)
   }
   return math.Abs(y-x)
}
func gravity(g,r float64) float64{
   e  := g/math.Pow(r,2)
   return  e
}
func mass(mofbase,massOfMotor,weightofpeople float64) float64{
   e:= mofbase*massOfMotor*weightofpeople
   return e
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

func main() {
 e:= ele{}
 e.name= "Four-locking mechanism"
 e.finalFloor=1000.5
 e.bottomFloor=0.0
   g:=0.0
 for i:= 6371.0; i<=42157.0; i++{
    g=gravity(9.8,i)
   fmt.Println(g)
    g+=g
}
   fmt.Println(g)

}
