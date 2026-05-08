package intermediate

import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
}

type rect struct{
	width,height float64
}
type circle struct{
	radius float64
}

func(r rect) area() float64{
	return r.height * r.width
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64{
	return 2 *(r.height + r.width)
}

func (c circle) perim() float64{
	return 2 *math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r := rect{width: 3,height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
	myPrinter(1,"john",45.9,true)

}
func printType(i interface{}){
	switch i.(type){
	case int:
		fmt.Println("type: int")
	case string:
		fmt.Println("type : string")
	default:
		fmt.Println("type: unknown")
	}
}

func myPrinter(i ...interface{}){
	for _,v:= rnage i {
		fmt.Println(v)
	}
}