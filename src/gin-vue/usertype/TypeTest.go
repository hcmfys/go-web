package main

import "fmt"

type Student struct {
	Name string
}

func t(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int ", i)
	}

}
func main() {

	dog := Student{}
	//_,_:=interface{}.(ds).(Student)
	//_,s:= interface{}.(dog).(Student)
	s, ok := interface{}(dog).(Student)
	fmt.Println(dog)
	fmt.Println(ok)
	fmt.Println(s)

	t(12)
}
