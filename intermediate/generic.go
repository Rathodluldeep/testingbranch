package intermediate

import (
	"fmt"
)

// func swap[T any](a,b T)(T ,T){
// 	return b,a
// }

type Stack[T any] struct {
	elements[]T 
}

func (s *Stack[T]) push(element T){
	s.elements=append(s.elements,element)
}

func (s *Stack[T]) pop() (T, bool){
	if len(s.elements)==0{
		var zero T
		return zero,false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element,true
}

func (s *Stack[T]) isEmpty() bool{
	return len(s.elements)==0
}
func (s *Stack[T]) PrintAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Print("Stack element: ")
	for _, element := range s.elements {
		fmt.Print(element, " ")
	}
	fmt.Println()
}
func main() {
	// x,y:= 1,2
	// x,y=swap(x,y)
	// fmt.Println(x, y)

	// x1,y1 := "hello","world"
	// x1,y1=swap(x1,y1)
	// fmt.Println(x1, y1)
	instack := Stack[int]{}
	instack.push(10)
	instack.push(20)
	instack.PrintAll()
	fmt.Println(instack.pop())
	fmt.Println(instack.pop())
	fmt.Println(instack.pop())
	instack.pop()
	instack.PrintAll() 


	stringstack := Stack[string]{}
	stringstack.push("hello")
	stringstack.push("world")
	stringstack.PrintAll()
	fmt.Println(stringstack.pop())
	fmt.Println(stringstack.pop())
	fmt.Println(stringstack.pop())
	stringstack.pop()
	stringstack.PrintAll() 
}