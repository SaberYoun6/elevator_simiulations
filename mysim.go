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
func force( m, g ,r float64)float64{
  velo := math.Sqrt(r*g/m)
  return  velo
}
func main() {
 e:= ele{}
 e.name= "Four-locking mechanism"
 e.finalFloor=1000.5
 r:= 35768.0
 d:= r+6731.0
 e.bottomFloor=0.0
   g:=0.0
 for i:= 0.0; i<=r; i++{
    g=gravity(9.8,i)
    g+=g
}
   dep:=0.0
 for q:= 6731.0;  q<=d; q++{
   dep=gravity(9.8,q)
   dep+=dep
}
   deep :=math.Pow(g,2)
   length :=math.Pow(dep,2)
   gd:=math.Sqrt(length+deep)
   fmt.Println(gd)
   w:=203000.0+185000.0+173000.0
   f:=force(w,g,r)
   forc:=force(w,gd,d)
   froce:=force(w,gd,6731.0)
   fmt.Printf("force from the distance of the whole stutucre with the the radius include the earth %f , The force from the radius with %f: just from ground to  the orbital platform %f\n",forc,froce,f)
   gp:= f/6300000000.0*r
   fmt.Println(g)
   fmt.Println(gp)

}
