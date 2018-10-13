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
func main() {
 e:= ele{}
 e.name= "Top-like zelda"
 e.bottomFloor =  0.0
 e.finalFloor = 1000.5
 i := isGoingUp(e.bottomFloor,e.finalFloor)
 n := isGoingDown(750.5,e.finalFloor)
 fmt.Printf("%+v",e)
 fmt.Println(i)
 fmt.Println(n)
 fmt.Println("hello W0rld")
}
